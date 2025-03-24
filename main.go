package main

import (
	"log"
	"os"

	"github.com/immanuel-254/potential-go/core/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sqlx.Open("sqlite3", os.Getenv("DB"))

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Enable foreign key support
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}

	database.DB = db

	defer func() {
		if closeError := db.Close(); closeError != nil {
			if err == nil {
				log.Fatalf("%s", closeError.Error())
				err = closeError
			}
		}
	}()

	goose.SetDialect("sqlite3")
}
