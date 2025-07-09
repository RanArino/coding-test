# Go ORM Testing Environment

This testing environment allows you to compare the performance and usage of three popular Go ORMs: GORM, Ent, and SQLC.

## 🚀 Quick Start

To run all tests and benchmarks from within the `code/tests` directory:

```bash
./run_tests.sh
```

This script will:
1. Install all dependencies for the test runner.
2. Test each ORM sample individually.
3. Run performance benchmarks comparing all three ORMs.
4. Display results with the fastest ORM highlighted.

## 🐳 Running with Docker

Alternatively, you can run the entire benchmark suite using Docker Compose. This is the recommended way to ensure a consistent testing environment.

From the `code/tests` directory, run:
```bash
docker compose up --build
```

This command will build the Docker image and run the `run_tests.sh` script inside a container. The benchmark typically takes less than a minute to run on modern hardware. For resource-constrained systems, you can adjust Docker's memory and CPU allocation.

## 📁 Project Structure

```
go_orms_research/
└── code/
    ├── samples/
    │   ├── gorm/          # GORM example
    │   ├── ent/           # Ent example  
    │   └── sqlc/          # SQLC example
    └── tests/
        ├── test_runner.go # Benchmark script
        ├── run_tests.sh   # Main test script
        ├── go.mod         # Test dependencies
        └── README.md      # This file
```

## 🧪 Individual ORM Testing

The `run_tests.sh` script handles testing all samples. To run them manually, navigate to each directory (e.g., `cd ../samples/gorm`) and execute `go run main.go`.

## ⚡ Performance Benchmarks

The `run_tests.sh` script also runs the benchmarks. To run them separately:

```bash
go run test_runner.go
```

## 🔧 Code Generation

### Ent
To regenerate Ent code after schema changes:
```bash
cd ../samples/ent && go generate ./...
```

### SQLC
To regenerate SQLC code after SQL changes:
```bash
cd ../samples/sqlc && sqlc generate
```

## 📊 Expected Results

Based on the article's research, you should see performance results similar to:

1. **SQLC** - Fastest (raw SQL performance)
2. **Ent** - Very good performance with type safety
3. **GORM** - Good performance with highest ease of use

## 🛠️ Requirements

### Native Execution
- Go 1.22+
- SQLite (handled by go-sqlite3 driver)
- SQLC tool (installed automatically)

### Docker Execution
- Docker and Docker Compose

## 📖 ORM Comparison

| Feature | GORM | Ent | SQLC |
|---------|------|-----|------|
| **Philosophy** | Feature-rich | Code generation | SQL-first |
| **Type Safety** | Runtime | Compile-time | Compile-time |
| **Ease of Use** | Easy | Medium | Medium |
| **Flexibility** | High | Medium | High |
| **Performance** | Good | Very Good | Very Good |

## 🎯 Use Cases

- **GORM**: Rapid prototyping, small projects, teams preferring ORM convenience
- **Ent**: Medium to large applications requiring type safety and good performance
- **SQLC**: Performance-critical applications, teams comfortable with SQL

## 🔍 Code Examples

### GORM Style
```go
// Create
db.Create(&Product{Code: "D42", Price: 100})

// Read
var product Product
db.First(&product, "code = ?", "D42")

// Update
db.Model(&product).Update("Price", 200)

// Delete
db.Delete(&product, 1)
```

### Ent Style
```go
// Create
user, err := client.User.
    Create().
    SetAge(30).
    SetName("a8m").
    Save(ctx)
```

### SQLC Style
```go
// First define SQL queries in .sql files
// Then use generated functions
user, err := queries.GetUser(ctx, 1)
```

## 📝 Article Reference

This testing environment supports the article "Go ORM徹底比較：GORM vs Ent vs SQLC (2025年)" which provides detailed analysis and recommendations for choosing the right ORM for your Go project. 