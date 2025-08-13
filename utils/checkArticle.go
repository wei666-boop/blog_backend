package utils

import (
	"demo/common"
	"demo/model"
	"github.com/gin-gonic/gin"
)

func CheckArticle(ctx *gin.Context) string {
	var articleID string
	articleID = ctx.Param("article_id")
	if articleID == "" {
		return ""
	}
	var article model.Article
	common.GetDB().First(&article, articleID)
	if article.Title == "" {
		return ""
	}
	return articleID
}
