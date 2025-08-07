package controller

import (
	"demo/common"
	"demo/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Login(ctx *gin.Context) {
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	var user model.User
	common.GetDB().Where("telephone = ?", telephone).First(&user)
	//验证手机号
	if user.ID == 0 {
		ctx.JSON(422, common.HttpResponse(422, user.ID, "验证失败"))
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(400, common.HttpResponse(400, nil, "验证失败"))
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(500, common.HttpResponse(500, err, "系统出错"))
		log.Printf("toke err:%#v\n", err)
		return
	}
	ctx.JSON(200, common.HttpResponse(200, token, "登录成功"))
}
