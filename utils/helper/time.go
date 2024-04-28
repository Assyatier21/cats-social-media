package helper

import (
	"log"
	"time"

	"github.com/backendmagang/project-1/models/entity"
)

var (
	jakartaLoc, _  = time.LoadLocation("Asia/Jakarta")
	TimeNowJakarta = time.Now().In(jakartaLoc)
)

func FormatTimeArticleResponse(article entity.ArticleResponse) entity.ArticleResponse {
	article.CreatedAt = FormattedTime(article.CreatedAt)
	article.UpdatedAt = FormattedTime(article.UpdatedAt)
	return article
}

func FormattedTime(ts string) string {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		log.Println(err)
		return ""
	}

	return t.Format("2006-01-02 15:04:05")
}
