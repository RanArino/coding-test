package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/gemini/sqlboiler-sample/models"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	boil.SetDB(db)

	// Insert a new user
	newUser := models.User{
		Name: "SQL Boiler",
	}
	err = newUser.Insert(ctx, db, boil.Infer())
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}
	fmt.Println("Inserted new user with ID:", newUser.ID)

	// Query for a user
	user, err := models.Users(models.UserWhere.Name.EQ("SQL Boiler")).One(ctx, db)
	if err != nil {
		log.Fatalf("failed to find user: %v", err)
	}

	fmt.Printf("Found user: ID=%d, Name=%s\n", user.ID.Int64, user.Name)
}
