package articledto

type ArticleResponse struct {
	ID     int    `json:"id"`
	Author string `json:"author" form:"author" validate:"required"`
	Title  string `json:"title" form:"title" validate:"required"`
	Body   string `json:"body" form:"body" validate:"required"`
}