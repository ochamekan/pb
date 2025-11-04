package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/visualect/pb/internal/handlers"
	"github.com/visualect/pb/internal/repo"
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

	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatalf("failed to conect to databse: %s\n", err)
	} else {
		log.Println("successfully connected to database")
	}

	articlesRepo := repo.New(dbpool)
	handler := handlers.New(articlesRepo)

	http.HandleFunc("GET /v1/getArticles", handler.GetArticles)
	http.HandleFunc("GET /v1/getArticle/{id}", handler.GetArticle)
	http.HandleFunc("POST /v1/createArticle", handler.CreateArticle)
	// http.HandleFunc("DELETE /v1/deleteArticle/{id}", deleteArticle)
	// http.HandleFunc("PATCH /v1/updateArticle/{id}", updateArticle)

	log.Println("running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
