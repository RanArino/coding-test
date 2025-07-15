# Ent Todo App - Educational Example

This is a simple todo application built with [Ent](https://entgo.io/) ORM to demonstrate proper Ent schema creation and usage.

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

## How to Run

1. **Backend:**
   ```bash
   cd backend
   go mod tidy
   go build -o todo-app .
   ./todo-app
   ```

2. **Frontend:**
   Open `frontend/index.html` in your browser

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