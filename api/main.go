package main

import (
	"fmt"
	"log"
	"net/http"

	"finance-helper/api/internal/db"
	"finance-helper/api/internal/handlers"
	"finance-helper/api/internal/migrations"
)

func main() {

	if err := db.Init(); err!= nil {
		log.Fatal("Database initialization failed:", err)
	}

	defer db.Close()

	if err := migrations.Run(db.GetDB(), "./migrations/sql"); err != nil {
		log.Fatal("Migration failed:", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)

	fmt.Println("Server is listening at :8080")

	http.ListenAndServe(":8080", mux)

}

