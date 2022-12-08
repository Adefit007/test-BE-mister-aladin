package repositories

import (
	"article/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindArticles() ([]models.Article, error)
	GetArticles(ID int) (models.Article, error)
	GetArticlesByAuthor(author string) ([]models.Article, error)
	GetArticlesByKeyword(keyword string) ([]models.Article, error)
	CreateArticles(articles models.Article)(models.Article, error)
	UpdateArticles(articles models.Article)(models.Article, error)
	DeleteArticles(articles models.Article)(models.Article, error)
}

type repository struct{
	db *gorm.DB
}

func RepositoryArticle(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArticles()([]models.Article,error)  {
	var articles []models.Article
	err := r.db.Find(&articles).Error

	return articles, err
}

func (r *repository) GetArticles(ID int) (models.Article,error)  {
	var articles models.Article
	err := r.db.First(&articles, ID).Error

	return articles,err
}

func (r *repository) GetArticlesByAuthor(author string) ([]models.Article,error)  {
	var articles []models.Article
	err := r.db.Where("author LIKE ?", "%"+author+"%").Find(&articles).Error

	return articles,err
}

func (r *repository) GetArticlesByKeyword(keyword string) ([]models.Article,error)  {
	var articles []models.Article
	err := r.db.Where("body LIKE ? || title LIKE ?", "%"+keyword+"%","%"+keyword+"%").Find(&articles).Error

	return articles,err
}

func (r *repository) CreateArticles(articles models.Article) (models.Article,error){
	err := r.db.Create(&articles).Error
	
	return articles, err
}

func (r *repository) UpdateArticles(articles models.Article) (models.Article, error)  {
	err := r.db.Save(&articles).Error

	return articles, err
}

func (r *repository) DeleteArticles(articles models.Article) (models.Article, error){
	err := r.db.Delete(&articles).Error

	return articles,err
}