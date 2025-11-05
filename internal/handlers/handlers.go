package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/visualect/pb/internal/repo"
)

type ArticlesHandler struct {
	repo repo.ArticlesRepository
}

func New(repo repo.ArticlesRepository) *ArticlesHandler {
	return &ArticlesHandler{repo}
}

func (a *ArticlesHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := a.repo.GetArticles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *ArticlesHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	article, err := a.repo.GetArticle(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *ArticlesHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var b repo.CreateArticleRequest
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if strings.Trim(b.Body, " ") == "" {
		http.Error(w, "body is required", http.StatusBadRequest)
		return
	}

	if strings.Trim(b.Title, " ") == "" {
		http.Error(w, "title is required", http.StatusBadRequest)
		return
	}

	article, err := a.repo.CreateArticle(r.Context(), b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *ArticlesHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = a.repo.DeleteArticle(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)
}
