package routes

import (
	"article/handlers"
	"article/repositories"

	"article/pkg/mysql"

	"github.com/gorilla/mux"
)

func ArticleRoutes(r *mux.Router){
	ArticleRepository := repositories.RepositoryArticle(mysql.DB)
	h := handlers.HandlerArticle(ArticleRepository)

	r.HandleFunc("/articles", h.FindArticles).Methods("GET")
	r.HandleFunc("/articles", h.CreateArticles).Methods("POST")
	r.HandleFunc("/articles/{id}", h.GetArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", h.UpdateArticles).Methods("PUT")
	r.HandleFunc("/articles/{id}", h.DeleteArticles).Methods("DELETE")
}