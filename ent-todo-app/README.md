# Ent Todo App - Educational Example

This is a simple todo application built with [Ent](https://entgo.io/) ORM to demonstrate proper Ent schema creation and usage.

## 📁 Getting This Code

### Option 1: Clone Only This Folder (Recommended)

If you only want this specific example without downloading the entire repository:

```bash
# 1. Clone only this folder using sparse checkout (Git 2.25+)
git clone --filter=blob:none --sparse https://github.com/RanArino/coding-test.git
cd coding-test
git sparse-checkout init --cone
git sparse-checkout set ent-todo-app

# 2. Navigate to the project directory
cd ent-todo-app
```

### Option 2: Clone Entire Repository

```bash
git clone https://github.com/RanArino/coding-test.git
cd coding-test/coding-test/ent-todo-app
```

**Why use Option 1?**
- This repository contains multiple independent article examples
- Saves bandwidth and storage by downloading only what you need
- Faster clone time

## Project Structure

```
ent-todo-app/
├── frontend/
│   └── index.html          # Simple HTML frontend
└── backend/
    ├── ent/                # Generated Ent code (DO NOT EDIT MANUALLY)
    │   ├── schema/         # Schema definitions
    │   │   ├── todo.go     # Todo entity schema
    │   │   └── user.go     # User entity schema
    │   ├── client.go       # Database client
    │   ├── ent.go          # Main Ent types
    │   ├── migrate/        # Migration utilities
    │   ├── todo/           # Todo entity operations
    │   ├── user/           # User entity operations
    │   └── ...             # Other generated files
    ├── go.mod
    ├── go.sum
    └── main.go             # Application entry point
```

## Key Educational Points

### 1. Ent Directory Structure
- **`ent/`** - Contains ALL generated code from Ent
- **`ent/schema/`** - Your entity schema definitions (this is what you edit)
- **`ent/client.go`** - Database client for operations
- **`ent/migrate/`** - Database migration utilities
- **`ent/{entity}/`** - Generated entity-specific operations

### 2. Schema Definition
Schema files in `ent/schema/` define your entities:
- `todo.go` - Defines Todo entity with fields and edges
- `user.go` - Defines User entity with fields and edges

### 3. Generated Code
Ent generates all the CRUD operations, query builders, and type-safe code based on your schema definitions.

## How to Run and Test CRUD Operations

### 1. Start the Backend Server

Open a terminal and navigate to the backend directory:

```bash
cd backend
go mod tidy
go run main.go
```

The server will start on `http://localhost:8080` and you'll see logs like:
```
[INFO] 🚀 Ent ToDo Application Starting
[SCHEMA] 📋 User Schema: Fields(name:string, email:string) + Edges(todos)
[SCHEMA] 📝 Todo Schema: Fields(text:string, status:enum, owner_id:int) + Edges(owner)
[INFO] 🗄️ Creating database schema from Ent definitions...
[INFO] ✅ Database tables created: users, todos
[INFO] 🌐 Server started on http://localhost:8080
```

### 2. Access the Frontend

The frontend is automatically served by the backend at: **http://localhost:8080**

Simply open your browser and go to `http://localhost:8080` to access the web interface.

### 3. Perform CRUD Operations

#### Via Web Interface:
- Open `http://localhost:8080` in your browser
- Use the forms to create users and todos
- View the real-time logs to see Ent ORM operations

#### Via API Endpoints (for testing):

**Create a User:**
```bash
curl -X POST -d "name=John&email=john@example.com" http://localhost:8080/api/users
```

**List Users:**
```bash
curl http://localhost:8080/api/users
```

**Create a Todo:**
```bash
curl -X POST -d "text=Learn Ent ORM&status=PENDING&user_id=1" http://localhost:8080/api/todos
```

**List Todos:**
```bash
curl http://localhost:8080/api/todos
```

**Valid Todo Status Values:** `PENDING`, `IN_PROGRESS`, `COMPLETED`

## Development Workflow

1. **Define/Modify Schema:** Edit files in `ent/schema/`
2. **Generate Code:** Run `go generate ./ent`
3. **Build & Run:** `go build` and run your application

## Important Notes

- **Never edit generated code** in `ent/` directory (except `ent/schema/`)
- **Always regenerate** after schema changes
- **The `ent/` directory is crucial** - it contains all ORM functionality
- **Schema files** are the source of truth for your data model

This structure demonstrates proper Ent usage for educational purposes, showing how schema definitions generate type-safe, feature-rich database operations. 