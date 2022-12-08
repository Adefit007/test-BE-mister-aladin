package handlers

import (
	articledto "article/dto/article"
	dto "article/dto/result"
	"article/models"
	"article/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerArticle struct {
	ArticleRepository repositories.ArticleRepository
}

func HandlerArticle(ArticleRepository repositories.ArticleRepository) *handlerArticle {
	return &handlerArticle{ArticleRepository}
}

func (h *handlerArticle) FindArticles(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	articles, err := h.ArticleRepository.FindArticles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: articles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) GetArticles(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	articles, err := h.ArticleRepository.GetArticles(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: articles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) CreateArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")

	request := articledto.ArticleRequest{
		Author: r.FormValue("author"),
		Title: r.FormValue("title"),
		Body: r.FormValue("body"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	
	articles := models.Article{
		Author: request.Author,
		Title: request.Title,
		Body: request.Body,
	}

	articles, err = h.ArticleRepository.CreateArticles(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	articles, _ =h.ArticleRepository.GetArticles(articles.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK,Data: articles}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) UpdateArticles(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	request := articledto.UpdateArticles{
		Author: r.FormValue("author"),
		Title: r.FormValue("title"),
		Body: r.FormValue("body"),
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	articles, err := h.ArticleRepository.GetArticles(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if (request.Author) != "" {
		articles.Author = request.Author
	}

	if (request.Title) != "" {
		articles.Title = request.Title
	}

	if (request.Body) != "" {
		articles.Body = request.Body
	}

	articles, err = h.ArticleRepository.UpdateArticles(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseArticles(articles)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerArticle) DeleteArticles(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	articles, err := h.ArticleRepository.GetArticles(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest,Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	data, err := h.ArticleRepository.DeleteArticles(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK,Data: data}
	json.NewEncoder(w).Encode(response)
}	

func convertResponseArticles(u models.Article) articledto.ArticleResponse  {
	return articledto.ArticleResponse{
		ID: u.ID,
		Author: u.Author,
		Title: u.Title,
		Body: u.Body,
	}
}