package main

import (
	"fmt"
	"log"
	"net/http"

	"finance-helper/api/internal/db"
	"finance-helper/api/internal/handlers"
	"finance-helper/api/internal/migrations"
	"finance-helper/api/internal/repository"
)

func main() {
	// Initialize database
	if err := db.Init(); err != nil {
		log.Fatal("Database initialization failed:", err)
	}
	defer db.Close()

	// Run migrations
	if err := migrations.Run(db.GetDB(), "./migrations/sql"); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.GetDB())

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userRepo)

	// Setup routes
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			userHandler.Create(w, r)
		} else if r.Method == http.MethodGet {
			userHandler.GetAll(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/users/by-id", userHandler.GetByID)

	fmt.Println("Server is listening at :8080")
	http.ListenAndServe(":8080", mux)
}