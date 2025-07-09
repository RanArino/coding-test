package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"sqlc.dev/example/db"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	// create tables
	_, err = conn.ExecContext(ctx, `CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT NOT NULL);`)
	if err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	// insert user
	_, err = queries.CreateUser(ctx, db.CreateUserParams{
		ID:   1,
		Name: "a8m",
	})
	if err != nil {
		log.Fatal(err)
	}

	user, err := queries.GetUser(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user)
}
