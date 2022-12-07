package articledto

type ArticleRequest struct {
	Author string `json:"author" gorm:"type:varchar(255)" validate:"required"`
	Title  string `json:"title" gorm:"type:varchar(255)" validate:"required"`
	Body   string `json:"body" gorm:"type:varchar(255)" validate:"required"`
}

type UpdateArticles struct {
	Author string `json:"author" form:"author"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}