package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/visualect/pb/internal/handlers"
	"github.com/visualect/pb/internal/handlers/middleware"
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
	h := handlers.New(articlesRepo)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/getArticles", h.GetArticles)
	mux.HandleFunc("GET /v1/getArticle/{id}", h.GetArticle)
	mux.HandleFunc("POST /v1/createArticle", middleware.AuthRequired(h.CreateArticle))
	mux.HandleFunc("DELETE /v1/deleteArticle/{id}", middleware.AuthRequired(h.DeleteArticle))
	mux.HandleFunc("PATCH /v1/updateArticle/{id}", middleware.AuthRequired(h.UpdateArticle))

	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		MaxAge:         86400,
	})
	handler := c.Handler(mux)

	log.Println("running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
