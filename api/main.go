package main

import (
	"fmt"
	"log"
	"net/http"

	"finance-helper/api/internal/db"
	"finance-helper/api/internal/handlers"
)

func main() {

	if err := db.Init(); err!= nil {
		log.Fatal("Database initialization failed:", err)
	}

	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)

	fmt.Println("Server is listening at :8080")

	http.ListenAndServe(":8080", mux)

}

