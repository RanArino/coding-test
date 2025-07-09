package main

import (
	"context"
	"log"

	"entgo.io/example/ent"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Create a new user.
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
}
