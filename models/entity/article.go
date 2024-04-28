package entity

type Article struct {
	ID          string   `json:"id" form:"id"`
	Title       string   `json:"title" form:"title"`
	Slug        string   `json:"slug" form:"slug"`
	HTMLContent string   `json:"html_content" form:"html_content"`
	CategoryIDs []int    `json:"category_id" form:"category_id"`
	Metadata    Metadata `json:"metadata" form:"metadata"`
	CreatedAt   string   `json:"created_at" form:"created_at"`
	UpdatedAt   string   `json:"updated_at" form:"updated_at"`
}
type Metadata struct {
	Title       string   `json:"meta_title"`
	Description string   `json:"meta_description"`
	Author      string   `json:"meta_author"`
	Keywords    []string `json:"meta_keywords"`
	Robots      []string `json:"meta_robots"`
}

type ArticleResponse struct {
	ID           string             `json:"id"`
	Title        string             `json:"title"`
	Slug         string             `json:"slug"`
	HTMLContent  string             `json:"html_content"`
	Metadata     string             `json:"metadata"`
	CreatedAt    string             `json:"created_at"`
	UpdatedAt    string             `json:"updated_at"`
	CategoryList []CategoryResponse `json:"categories"`
}

type GetArticlesRequest struct {
	Limit       int    `param:"limit" json:"limit" form:"limit"`
	Offset      int    `param:"offset" json:"offset" form:"offset"`
	SortBy      string `param:"sort_by" json:"sort_by" form:"sort_by"`
	OrderBy     string `param:"order_by" json:"order_by" form:"order_by"`
	OrderByBool bool
}

type GetArticleDetailsRequest struct {
	ID string `json:"id" form:"id" param:"id" validate:"required"`
}

type InsertArticleRequest struct {
	ID          string `json:"id" form:"id"`
	Title       string `json:"title" form:"title" validate:"required"`
	Slug        string `json:"slug" form:"slug" validate:"required"`
	HTMLContent string `json:"htmlcontent" form:"html_content" validate:"required"`
	CategoryIDs []int  `json:"category_id" form:"category_id" validate:"required"`
	Metadata    string `json:"metadata" form:"metadata" validate:"required"`
	CreatedAt   string `json:"created_at" form:"created_at"`
	UpdatedAt   string `json:"updated_at" form:"updated_at"`
}

type UpdateArticleRequest struct {
	ID          string `json:"id" form:"id"`
	Title       string `json:"title" form:"title"`
	Slug        string `json:"slug" form:"slug"`
	HTMLContent string `json:"htmlcontent" form:"html_content"`
	CategoryIDs []int  `json:"category_id" form:"category_id"`
	Metadata    string `json:"metadata" form:"metadata"`
	CreatedAt   string `json:"created_at" form:"created_at"`
	UpdatedAt   string `json:"updated_at" form:"updated_at"`
}

type DeleteArticleRequest struct {
	ID string `json:"id" form:"id" param:"id" validate:"required"`
}
