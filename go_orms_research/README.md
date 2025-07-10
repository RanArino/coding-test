# Go ORM Comparison: GORM vs Ent vs SQLC

A comprehensive testing environment to compare the performance and usage of three popular Go ORMs: **GORM**, **Ent**, and **SQLC**.

## 🚀 Quick Start

### Prerequisites

**For Native Execution:**
- Go 1.22 or higher
- SQLite (handled automatically by go-sqlite3 driver)

**For Docker Execution:**
- Docker and Docker Compose

## 🐳 Running Tests with Docker (Recommended)

Docker provides a consistent testing environment and is the recommended approach.

```bash
# Navigate to the tests directory
cd tests

# Run the complete test suite with Docker
docker compose up --build
```

This will:
1. Build a containerized environment with Go 1.22 and all dependencies
2. Test each ORM sample individually
3. Run performance benchmarks comparing all three ORMs
4. Display results with the fastest ORM highlighted

### Clean up Docker containers
```bash
docker compose down
```

## 💻 Running Tests Natively

For local development and testing without Docker:

```bash
# Navigate to the tests directory
cd tests

# Install dependencies
go mod tidy

# Run the comprehensive test script
chmod +x run_tests.sh
./run_tests.sh
```

### Testing Individual ORM Samples

You can also test each ORM sample individually:

```bash
# Test GORM sample
cd samples/gorm
go mod tidy
go run main.go

# Test Ent sample
cd ../ent
go mod tidy
go run main.go

# Test SQLC sample
cd ../sqlc
go mod tidy
go run main.go
```

### Running Performance Benchmarks Only

```bash
cd tests
go run test_runner.go
```

## 📁 Project Structure

```
go_orms_research/
├── README.md              # This file
├── samples/
│   ├── gorm/              # GORM implementation example
│   │   ├── main.go
│   │   ├── go.mod
│   │   └── go.sum
│   ├── ent/               # Ent implementation example
│   │   ├── main.go
│   │   ├── generate.go
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── ent/           # Generated Ent code
│   └── sqlc/              # SQLC implementation example
│       ├── main.go
│       ├── schema.sql
│       ├── query.sql
│       ├── sqlc.yaml
│       ├── go.mod
│       ├── go.sum
│       └── db/            # Generated SQLC code
└── tests/
    ├── README.md          # Detailed testing documentation
    ├── test_runner.go     # Performance benchmark script
    ├── run_tests.sh       # Comprehensive test script
    ├── docker-compose.yml # Docker configuration
    ├── Dockerfile         # Docker image definition
    ├── go.mod
    └── go.sum
```

## 📊 Expected Performance Results

Based on benchmark testing, you should see performance results similar to:

| Rank | ORM | Typical Performance | Characteristics |
|------|-----|---------------------|----------------|
| 🥇 1st | **SQLC** | ~10-15ms | Raw SQL performance, compile-time safety |
| 🥈 2nd | **Ent** | ~10-16ms | Type-safe, good performance, rich features |
| 🥉 3rd | **GORM** | ~60-85ms | Feature-rich, easy to use, some overhead |

*Results based on 1000 CRUD operations in containerized environment*

## 🔧 Code Generation

### Regenerating Ent Code
If you modify the Ent schema:
```bash
cd samples/ent
go generate ./...
```

### Regenerating SQLC Code
If you modify SQL queries:
```bash
cd samples/sqlc
sqlc generate
```

## 🧪 What Each Test Does

### Individual ORM Samples
Each sample demonstrates basic CRUD operations:
- **Create**: Insert new records
- **Read**: Query existing data
- **Update**: Modify existing records
- **Delete**: Remove records

### Performance Benchmarks
The benchmark test performs 1000 CRUD operations for each ORM and measures:
- Total execution time
- Relative performance comparison
- Winner determination

## 📖 ORM Comparison Summary

| Feature | GORM | Ent | SQLC |
|---------|------|-----|------|
| **Philosophy** | Feature-rich ORM | Code generation + type safety | SQL-first approach |
| **Type Safety** | Runtime | Compile-time | Compile-time |
| **Learning Curve** | Easy | Medium | Medium |
| **Performance** | Good | Very Good | Excellent |
| **Flexibility** | High | Medium | High |
| **Code Generation** | No | Yes | Yes |

## 🎯 Use Case Recommendations

- **Choose GORM if**: You want rapid prototyping, have a small project, or prefer ORM convenience
- **Choose Ent if**: You need type safety, have medium to large applications, and want good performance
- **Choose SQLC if**: Performance is critical, you're comfortable with SQL, or need maximum control

## 🚨 Troubleshooting

### Common Issues

**Docker build fails:**
- Ensure Docker is running
- Try: `docker system prune` to clean up

**Native execution fails:**
- Check Go version: `go version` (needs 1.22+)
- Install dependencies: `go mod tidy`

**Permission denied on run_tests.sh:**
```bash
chmod +x run_tests.sh
```

**CGO errors:**
- Ensure you have a C compiler installed
- On macOS: `xcode-select --install`

### Getting Help

1. Check that all dependencies are properly installed
2. Verify your Go version meets requirements
3. For Docker issues, ensure Docker Desktop is running
4. Review the detailed logs in `tests/README.md`

## 📝 License

This project is for educational and comparison purposes. Each ORM has its own license terms. 