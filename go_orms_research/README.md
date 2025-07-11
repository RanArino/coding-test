# Go ORM Comparison: GORM vs Ent vs SQLC vs SQLBoiler

A comprehensive testing environment to compare the performance and usage of four popular Go ORMs: **GORM**, **Ent**, **SQLC**, and **SQLBoiler**.

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
3. Run performance benchmarks comparing all four ORMs
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
│   ├── ent/               # Ent implementation example
│   ├── sqlc/              # SQLC implementation example
│   └── sqlboiler/         # SQLBoiler implementation example
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
| 🥇 1st | **SQLC** | ~15ms | Raw SQL performance, compile-time safety |
| 🥈 2nd | **Ent** | ~16ms | Type-safe, good performance, rich features |
| 🥉 3rd | **SQLBoiler** | ~17ms | Database-first, generated models, good performance |
| 4th | **GORM** | ~80-90ms | Feature-rich, easy to use, some overhead |

*Results based on 1000 CRUD operations in the containerized Docker environment.*

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

### Regenerating SQLBoiler Code
If you modify the database schema:
```bash
cd samples/sqlboiler
# Ensure the database file (test.db) is up to date with the schema
sqlboiler sqlite3
```

## 🏁 Conclusion & Performance Analysis

This testing environment provides a robust comparison of four major Go ORMs. The benchmarks reveal a clear performance hierarchy and highlight how the execution environment can influence results.

### Benchmark Results Summary

The final benchmark results for 1,000 CRUD operations were as follows:

| ORM         | Local (macOS, Go 1.22) | Docker (Linux, Go 1.22) |
|-------------|------------------------|-------------------------|
| **SQLBoiler** | **~9.1ms (🥇 1st)**    | ~16.8ms (3rd)           |
| **SQLC**      | ~10.6ms (2nd)          | **~15.1ms (🥇 1st)**    |
| **Ent**       | ~11.7ms (3rd)          | ~15.6ms (2nd)           |
| **GORM**      | ~57.5ms (4th)          | ~82.6ms (4th)           |

### Analysis: Why Did the Rankings Change?

An interesting observation is the performance shift between the local and Docker environments. Locally, SQLBoiler was the fastest, but in the Docker container, SQLC took the lead. This discrepancy is an excellent example of why testing in a production-like environment is critical.

The potential reasons for this behavior include:

1.  **CGO Toolchain Differences**: The `mattn/go-sqlite3` driver uses CGO to interface with SQLite's C library. The performance of these Go-to-C calls is affected by the underlying C compiler and standard library.
    *   **Local (macOS)**: Uses the Clang/LLVM toolchain.
    *   **Docker (Alpine Linux)**: Uses GCC and the `musl` C standard library.
    The interaction between Go and C in the Alpine/`musl` environment has a slightly different overhead profile, which was enough to impact SQLBoiler's ranking when margins are measured in single-digit milliseconds.

2.  **Kernel and System Call Variations**: Docker runs on a Linux kernel (via a lightweight VM on macOS). System calls for file I/O, memory management, and process scheduling differ between the macOS (Darwin) and Linux kernels. These subtle differences can influence the performance of I/O-bound or CPU-intensive tasks.

3.  **Virtualization Overhead**: The virtualization layer that Docker uses on macOS introduces a small but constant performance cost. This, combined with Docker's overlay filesystem, creates an environment with different characteristics than running natively on APFS.

In conclusion, when performance margins are extremely tight, even minor environmental factors can be enough to reorder the rankings. The containerized results should be considered more representative of a typical production deployment.

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
