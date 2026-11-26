package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/pratyoos/primestay-backend/internal/platform/config"
	"github.com/pratyoos/primestay-backend/internal/platform/database"
)

func main() {
	_ = godotenv.Load()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	pool, err := database.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer pool.Close()
	fmt.Println("connected to database")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status": "ok"}`)
	})

	fmt.Println("server running on port", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}