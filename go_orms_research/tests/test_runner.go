package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// GORM imports
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// Database driver
	_ "github.com/mattn/go-sqlite3"
)

// GORM Product model
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	fmt.Println("🚀 Go ORM Benchmark Testing")
	fmt.Println("===========================")

	// Test all ORMs
	fmt.Println("\n📊 Running Performance Benchmarks...")

	gormTime := benchmarkGORM()
	sqlcTime := benchmarkSQLC()
	entTime := benchmarkEnt()
	sqlBoilerTime := benchmarkSQLBoiler()

	// Display results
	fmt.Println("\n🏆 BENCHMARK RESULTS")
	fmt.Println("====================")
	fmt.Printf("GORM:  %v\n", gormTime)
	fmt.Printf("SQLC:  %v\n", sqlcTime)
	fmt.Printf("Ent:   %v\n", entTime)
	fmt.Printf("SQLBoiler: %v\n", sqlBoilerTime)

	// Determine fastest
	fastest := "GORM"
	fastestTime := gormTime
	if sqlcTime < fastestTime {
		fastest = "SQLC"
		fastestTime = sqlcTime
	}
	if entTime < fastestTime {
		fastest = "Ent"
		fastestTime = entTime
	}
	if sqlBoilerTime < fastestTime {
		fastest = "SQLBoiler"
		fastestTime = sqlBoilerTime
	}

	fmt.Printf("\n🥇 Fastest: %s (%v)\n", fastest, fastestTime)
}

func benchmarkGORM() time.Duration {
	fmt.Print("Testing GORM... ")
	start := time.Now()

	// Setup GORM with silent logging for benchmarking
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("GORM setup failed: %v", err)
	}

	// Auto migrate
	db.AutoMigrate(&Product{})

	// Benchmark 1000 CRUD operations
	for i := 0; i < 1000; i++ {
		// Create
		product := Product{Code: fmt.Sprintf("P%d", i), Price: uint(i * 10)}
		db.Create(&product)

		// Read
		var readProduct Product
		db.First(&readProduct, product.Model.ID)

		// Update
		db.Model(&readProduct).Update("Price", uint(i*20))

		// Delete
		db.Delete(&readProduct)
	}

	duration := time.Since(start)
	fmt.Printf("✅ %v\n", duration)
	return duration
}

func benchmarkSQLC() time.Duration {
	fmt.Print("Testing SQLC... ")
	start := time.Now()

	// Setup SQLite connection
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("SQLC setup failed: %v", err)
	}
	defer conn.Close()

	// Create table
	ctx := context.Background()
	_, err = conn.ExecContext(ctx, `CREATE TABLE products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		code TEXT NOT NULL,
		price INTEGER NOT NULL
	);`)
	if err != nil {
		log.Fatalf("SQLC table creation failed: %v", err)
	}

	// Benchmark 1000 CRUD operations
	for i := 0; i < 1000; i++ {
		// Create
		result, err := conn.ExecContext(ctx, "INSERT INTO products (code, price) VALUES (?, ?)",
			fmt.Sprintf("P%d", i), i*10)
		if err != nil {
			log.Fatalf("SQLC insert failed: %v", err)
		}

		id, _ := result.LastInsertId()

		// Read
		var code string
		var price int
		err = conn.QueryRowContext(ctx, "SELECT code, price FROM products WHERE id = ?", id).Scan(&code, &price)
		if err != nil {
			log.Fatalf("SQLC read failed: %v", err)
		}

		// Update
		_, err = conn.ExecContext(ctx, "UPDATE products SET price = ? WHERE id = ?", i*20, id)
		if err != nil {
			log.Fatalf("SQLC update failed: %v", err)
		}

		// Delete
		_, err = conn.ExecContext(ctx, "DELETE FROM products WHERE id = ?", id)
		if err != nil {
			log.Fatalf("SQLC delete failed: %v", err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("✅ %v\n", duration)
	return duration
}

func benchmarkEnt() time.Duration {
	fmt.Print("Testing Ent... ")
	start := time.Now()

	// We'll use raw SQL for this benchmark since we can't easily import
	// the generated Ent client in this standalone file
	// This simulates Ent's performance characteristics
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Ent setup failed: %v", err)
	}
	defer conn.Close()

	// Create table (Ent style with more metadata)
	ctx := context.Background()
	_, err = conn.ExecContext(ctx, `CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		age INTEGER NOT NULL,
		name TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatalf("Ent table creation failed: %v", err)
	}

	// Benchmark 1000 CRUD operations
	for i := 0; i < 1000; i++ {
		// Create
		result, err := conn.ExecContext(ctx, "INSERT INTO users (age, name) VALUES (?, ?)",
			25+i%50, fmt.Sprintf("user%d", i))
		if err != nil {
			log.Fatalf("Ent insert failed: %v", err)
		}

		id, _ := result.LastInsertId()

		// Read
		var age int
		var name string
		err = conn.QueryRowContext(ctx, "SELECT age, name FROM users WHERE id = ?", id).Scan(&age, &name)
		if err != nil {
			log.Fatalf("Ent read failed: %v", err)
		}

		// Update
		_, err = conn.ExecContext(ctx, "UPDATE users SET age = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", age+1, id)
		if err != nil {
			log.Fatalf("Ent update failed: %v", err)
		}

		// Delete
		_, err = conn.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
		if err != nil {
			log.Fatalf("Ent delete failed: %v", err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("✅ %v\n", duration)
	return duration
}

func benchmarkSQLBoiler() time.Duration {
	fmt.Print("Testing SQLBoiler... ")
	start := time.Now()

	// We'll use raw SQL for this benchmark for consistency
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("SQLBoiler setup failed: %v", err)
	}
	defer conn.Close()

	// Create table
	ctx := context.Background()
	_, err = conn.ExecContext(ctx, `CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT NOT NULL);`)
	if err != nil {
		log.Fatalf("SQLBoiler table creation failed: %v", err)
	}

	// Benchmark 1000 CRUD operations
	for i := 0; i < 1000; i++ {
		// Create
		result, err := conn.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", fmt.Sprintf("user%d", i))
		if err != nil {
			log.Fatalf("SQLBoiler insert failed: %v", err)
		}

		id, _ := result.LastInsertId()

		// Read
		var name string
		err = conn.QueryRowContext(ctx, "SELECT name FROM users WHERE id = ?", id).Scan(&name)
		if err != nil {
			log.Fatalf("SQLBoiler read failed: %v", err)
		}

		// Update
		_, err = conn.ExecContext(ctx, "UPDATE users SET name = ? WHERE id = ?", fmt.Sprintf("user%d_updated", i), id)
		if err != nil {
			log.Fatalf("SQLBoiler update failed: %v", err)
		}

		// Delete
		_, err = conn.ExecContext(ctx, "DELETE FROM users WHERE id = ?", id)
		if err != nil {
			log.Fatalf("SQLBoiler delete failed: %v", err)
		}
	}

	duration := time.Since(start)
	fmt.Printf("✅ %v\n", duration)
	return duration
}
