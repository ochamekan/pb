package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/visualect/pb/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbpool.Close()

	http.HandleFunc("GET /v1/getArticles", getArticles)

	// GET getArticles (maybe filter or sort via search params)
	// GET getArticle
	// POST createArticle
	// DELETE deleteArticle
	// PATCH updateArticle

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article
	// db get articles

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
