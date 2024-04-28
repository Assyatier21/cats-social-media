package constant

const (
	OK      = "OK"
	SUCCESS = "success"
	FAILED  = "failed"

	SUCCESS_LOGIN         = "successfully login"
	SUCCESS_REGISTER_USER = "phone number succesfully registered"

	SUCCESS_GET_ARTICLES   = "success get list of articles"
	SUCCESS_INSERT_ARTICLE = "article inserted successfully"
	SUCCESS_UPDATE_ARTICLE = "article updated successfully"
	SUCCESS_DELETE_ARTICLE = "article deleted successfully"

	SUCCESS_GET_CATEGORY    = "success get category tree"
	SUCCESS_INSERT_CATEGORY = "category inserted successfully"
	SUCCESS_UPDATE_CATEGORY = "category updated successfully"
	SUCCESS_DELETE_CATEGORY = "category deleted successfully"

	FAILED_GET_ARTICLES            = "failed to get list of articles"
	FAILED_GET_ARTICLE_DETAILS     = "failed to get article details"
	FAILED_BUILD_ARTICLE_REQUEST   = "failed to build article request"
	FAILED_INSERT_ARTICLE_POSTGRES = "failed to insert article in postgres"
	FAILED_INSERT_ARTICLE_ELASTIC  = "failed to insert article in elasticsearch"
	FAILED_UPDATE_ARTICLE_POSTGRES = "failed to update article in postgres"
	FAILED_UPDATE_ARTICLE_ELASTIC  = "failed to update article in elasticsearch"
	FAILED_DELETE_ARTICLE_POSTGRES = "failed to delete article in postgres"
	FAILED_DELETE_ARTICLE_ELASTIC  = "failed to delete article in elasticsearch"

	FAILED_GET_CATEGORIES           = "failed to get category tree"
	FAILED_GET_CATEGORY_DETAILS     = "failed to get category details"
	FAILED_INSERT_CATEGORY_POSTGRES = "failed to insert category in postgres"
	FAILED_INSERT_CATEGORY_ELASTIC  = "failed to insert category in elasticsearch"
	FAILED_UPDATE_CATEGORY_POSTGRES = "failed to update category in postgres"
	FAILED_UPDATE_CATEGORY_ELASTIC  = "failed to update category in elasticsearch"
	FAILED_DELETE_CATEGORY_POSTGRES = "failed to delete category in postgres"
	FAILED_DELETE_CATEGORY_ELASTIC  = "failed to delete category in elasticsearch"

	FAILED_GET_CATS = "failed to get list of cats"
)
