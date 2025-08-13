package controller

import (
	"demo/common"
	"demo/model"
	"demo/utils"
	"github.com/gin-gonic/gin"
	"time"
)

func AddArticle(ctx *gin.Context) {

	var article model.Article
	var demo_article model.Article
	ctx.Bind(&article)

	ArticleID := article.ArticleID
	ArticleTitle := article.Title
	ArticleContent := article.Content
	ArticleAuthor := article.Author
	ArticleCreateTime := time.Now().Format("2006-01-02 15:04:05")

	common.GetDB().Where("article_id=?", ArticleID).First(&demo_article)
	if demo_article.ArticleID == ArticleID {
		ctx.JSON(422, common.HttpResponse(422, nil, "文章已存在,禁止重复发布"))
		return
	}

	if article.Title == "" || article.Content == "" {
		ctx.JSON(422, common.HttpResponse(422, nil, "文章标题或文章不能为空"))
		return
	}

	newArticle := model.Article{
		ArticleID:  ArticleID,
		Title:      ArticleTitle,
		Author:     ArticleAuthor,
		Content:    ArticleContent,
		CreateTime: ArticleCreateTime,
	}
	common.GetDB().Create(&newArticle)
	ctx.JSON(200, common.HttpResponse(200, nil, "文章发布成功"))
}

func DeleteArticle(ctx *gin.Context) {
	var articleID string
	var article model.Article
	articleID = utils.CheckArticle(ctx)
	if articleID == "" {
		ctx.JSON(422, common.HttpResponse(422, nil, "请检查是否传入参数或者文章存不存在"))
		return
	}
	if db := common.GetDB().Delete(&article); db.Error != nil {
		ctx.JSON(422, common.HttpResponse(422, db.Error, "删除过程发生错误"))
		return
	}
	ctx.JSON(200, common.HttpResponse(200, nil, "删除成功"))
}

func UpdateArticle(ctx *gin.Context) {
	var articleID string
	articleID = utils.CheckArticle(ctx)
	if articleID == "" {
		ctx.JSON(422, common.HttpResponse(422, nil, "请检查是否传入参数或者文章存不存在"))
		return
	}
	var article model.Article
	ctx.Bind(&article)

	var newArticle model.Article
	common.GetDB().First(&newArticle, articleID)
	newArticle.Title = article.Title
	newArticle.Content = article.Content
	newArticle.Author = article.Author
	newArticle.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	common.GetDB().Save(&newArticle)

	ctx.JSON(200, common.HttpResponse(200, nil, "文章更新成功"))
}
