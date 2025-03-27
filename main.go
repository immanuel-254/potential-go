package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/immanuel-254/potential-go/core/database"
	"github.com/immanuel-254/potential-go/core/views"
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

	// Apply all "up" migrations
	err = goose.Up(database.DB.DB, "core/migrations")
	if err != nil {
		log.Fatalf("Failed to auth apply migrations: %v", err)
	}

	server()
}

func server() {
	mux := http.NewServeMux()

	views.Routes(mux, views.UserViews)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", os.Getenv("PORT")), // Custom port
		//Handler:      internal.LoggingMiddleware(internal.Cors(internal.New(internal.ConfigDefault)(mux))), // Attach the mux as the handler
		Handler:      mux,
		ReadTimeout:  10 * time.Second, // Set read timeout
		WriteTimeout: 10 * time.Second, // Set write timeout
		IdleTimeout:  30 * time.Second, // Set idle timeout
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error starting server:", err)
	}
}
