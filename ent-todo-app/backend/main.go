package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"

	"ent-example/ent"
	"ent-example/ent/todo"
)

// LogEntry represents a structured log entry for frontend display
type LogEntry struct {
	Timestamp  string   `json:"Timestamp"`
	Level      string   `json:"Level"`
	Message    string   `json:"Message"`
	SQL        string   `json:"SQL,omitempty"`
	SQLParams  []string `json:"SQLParams,omitempty"`
	Operation  string   `json:"Operation,omitempty"`
	SchemaInfo string   `json:"SchemaInfo,omitempty"`
}

var (
	logEntries []LogEntry
	logMutex   sync.Mutex
	currentOp  string
)

type logWriter struct{}

func (w logWriter) Write(p []byte) (n int, err error) {
	logMutex.Lock()
	defer logMutex.Unlock()

	line := strings.TrimSpace(string(p))
	if line == "" {
		return len(p), nil
	}

	entry := LogEntry{
		Timestamp: time.Now().Format("15:04:05.000"),
		Level:     "SQL",
		Message:   line,
		Operation: currentOp,
	}

	// Parse SQL from Ent debug output
	if strings.Contains(line, "driver.Query:") || strings.Contains(line, "driver.Exec:") {
		parts := strings.Split(line, " ")
		for _, part := range parts {
			if strings.HasPrefix(part, "query=") {
				sqlQuery := strings.TrimPrefix(part, "query=")
				entry.SQL = sqlQuery
				entry.Message = fmt.Sprintf("Ent generated SQL: %s", sqlQuery)
				break
			}
		}
	}

	logEntries = append(logEntries, entry)
	fmt.Printf("[SQL] %s\n", line)
	return len(p), nil
}

func appendLog(level, message, operation, schemaInfo string) {
	logMutex.Lock()
	defer logMutex.Unlock()

	entry := LogEntry{
		Timestamp:  time.Now().Format("15:04:05.000"),
		Level:      level,
		Message:    message,
		Operation:  operation,
		SchemaInfo: schemaInfo,
	}

	logEntries = append(logEntries, entry)
	fmt.Printf("[%s] %s\n", level, message)
}

func startOperation(operation string) {
	currentOp = operation
	appendLog("OPERATION", "=== "+operation+" ===", operation, "")
}

func main() {
	log.SetOutput(logWriter{})
	log.SetFlags(0)

	appendLog("INFO", "🚀 Ent ToDo Application Starting", "STARTUP", "")
	appendLog("SCHEMA", "📋 User Schema: Fields(name:string, email:string) + Edges(todos)", "STARTUP", getUserSchemaInfo())
	appendLog("SCHEMA", "📝 Todo Schema: Fields(text:string, status:enum, owner_id:int) + Edges(owner)", "STARTUP", getTodoSchemaInfo())

	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=true")
	if err != nil {
		appendLog("ERROR", fmt.Sprintf("Failed to connect to database: %v", err), "STARTUP", "")
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	appendLog("INFO", "🗄️ Creating database schema from Ent definitions...", "STARTUP", "")
	if err := client.Schema.Create(context.Background()); err != nil {
		appendLog("ERROR", fmt.Sprintf("Failed to create schema: %v", err), "STARTUP", "")
		log.Fatalf("failed creating schema resources: %v", err)
	}
	appendLog("INFO", "✅ Database tables created: users, todos", "STARTUP", "")

	client = client.Debug()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../frontend/index.html")
	})
	http.HandleFunc("/api/users", handleUsers(client))
	http.HandleFunc("/api/todos", handleTodos(client))
	http.HandleFunc("/api/logs", handleLogs)
	http.HandleFunc("/api/schema", handleSchema)

	appendLog("INFO", "🌐 Server started on http://localhost:8080", "STARTUP", "")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleUsers(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			startOperation("Create User")
			name := r.FormValue("name")
			email := r.FormValue("email")

			appendLog("INFO", fmt.Sprintf("👤 Creating new user: name='%s', email='%s'", name, email), currentOp, getUserSchemaInfo())
			appendLog("ENT", "🔧 Ent: Building User.Create() mutation with SetName() and SetEmail()", currentOp, "")
			appendLog("ENT", "📊 Ent will generate INSERT SQL for 'users' table", currentOp, "")

			user, err := client.User.Create().SetName(name).SetEmail(email).Save(context.Background())
			if err != nil {
				appendLog("ERROR", fmt.Sprintf("❌ Failed to create user: %v", err), currentOp, "")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			appendLog("SUCCESS", fmt.Sprintf("✅ User created successfully! ID=%d, Name='%s'", user.ID, user.Name), currentOp, "")
			fmt.Fprintf(w, "Created User: %s (ID: %d)\n", user.Name, user.ID)

		case http.MethodGet:
			startOperation("List Users")
			appendLog("INFO", "📋 Fetching all users from database", currentOp, getUserSchemaInfo())
			appendLog("ENT", "🔧 Ent: Building User.Query().All() to fetch all user records", currentOp, "")
			appendLog("ENT", "📊 Ent will generate SELECT SQL for 'users' table", currentOp, "")

			users, err := client.User.Query().All(context.Background())
			if err != nil {
				appendLog("ERROR", fmt.Sprintf("❌ Failed to fetch users: %v", err), currentOp, "")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			appendLog("SUCCESS", fmt.Sprintf("✅ Retrieved %d users from database", len(users)), currentOp, "")
			fmt.Fprintln(w, "Users:")
			for _, u := range users {
				fmt.Fprintf(w, "- ID: %d, Name: %s, Email: %s\n", u.ID, u.Name, u.Email)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleTodos(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			startOperation("Create Todo")
			text := r.FormValue("text")
			userID := r.FormValue("user_id")
			status := r.FormValue("status")

			ownerID, err := strconv.Atoi(userID)
			if err != nil {
				appendLog("ERROR", fmt.Sprintf("❌ Invalid user ID: %s", userID), currentOp, "")
				http.Error(w, "Invalid User ID", http.StatusBadRequest)
				return
			}

			appendLog("INFO", fmt.Sprintf("📝 Creating todo: text='%s', status='%s', owner_id=%d", text, status, ownerID), currentOp, getTodoSchemaInfo())
			appendLog("ENT", "🔧 Ent: Building Todo.Create() with SetText(), SetStatus(), SetOwnerID()", currentOp, "")
			appendLog("ENT", "🔗 Ent: Establishing relationship between Todo and User via owner_id foreign key", currentOp, "")
			appendLog("ENT", "📊 Ent will generate INSERT SQL for 'todos' table with foreign key reference", currentOp, "")

			todoObj, err := client.Todo.Create().SetText(text).SetStatus(todo.Status(status)).SetOwnerID(ownerID).Save(context.Background())
			if err != nil {
				appendLog("ERROR", fmt.Sprintf("❌ Failed to create todo: %v", err), currentOp, "")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			appendLog("SUCCESS", fmt.Sprintf("✅ Todo created! ID=%d, Text='%s', Owner=%d", todoObj.ID, todoObj.Text, todoObj.OwnerID), currentOp, "")
			fmt.Fprintf(w, "Created Todo: %s (ID: %d, Owner: %d)\n", todoObj.Text, todoObj.ID, todoObj.OwnerID)

		case http.MethodGet:
			startOperation("List Todos")
			appendLog("INFO", "📋 Fetching all todos with their owners", currentOp, getTodoSchemaInfo())
			appendLog("ENT", "🔧 Ent: Building Todo.Query().WithOwner() for eager loading", currentOp, "")
			appendLog("ENT", "🔗 Ent: Will JOIN todos and users tables to load relationships", currentOp, "")
			appendLog("ENT", "📊 Ent will generate SELECT with JOIN to fetch todos + owner data", currentOp, "")

			todos, err := client.Todo.Query().WithOwner().All(context.Background())
			if err != nil {
				appendLog("ERROR", fmt.Sprintf("❌ Failed to fetch todos: %v", err), currentOp, "")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			appendLog("SUCCESS", fmt.Sprintf("✅ Retrieved %d todos with owner information", len(todos)), currentOp, "")
			fmt.Fprintln(w, "Todos:")
			for _, t := range todos {
				ownerName := "N/A"
				if t.Edges.Owner != nil {
					ownerName = t.Edges.Owner.Name
				}
				fmt.Fprintf(w, "- ID: %d, Text: %s, Status: %s, Owner: %s (ID: %d)\n", t.ID, t.Text, t.Status, ownerName, t.OwnerID)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func handleLogs(w http.ResponseWriter, r *http.Request) {
	logMutex.Lock()
	defer logMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(logEntries); err != nil {
		http.Error(w, "Failed to encode logs", http.StatusInternalServerError)
	}
}

func handleSchema(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	schema := map[string]interface{}{
		"user_schema":   getUserSchemaInfo(),
		"todo_schema":   getTodoSchemaInfo(),
		"relationships": "User has many Todos (one-to-many). Todo belongs to one User (many-to-one).",
	}
	json.NewEncoder(w).Encode(schema)
}

func getUserSchemaInfo() string {
	return "User{name:string(unique), email:string(unique)} → edges: todos[]"
}

func getTodoSchemaInfo() string {
	return "Todo{text:string, status:enum, owner_id:int} → edges: owner"
}
