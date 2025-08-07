package controller

import (
	"demo/common"
	"demo/model"
	"demo/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func RegisterUser(ctx *gin.Context) {
	//获取参数
	//使用map来获取参数
	//var requestMap = make(map[string]interface{})
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)

	//使用结构体来获取参数
	var requestUser = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	ctx.Bind(&requestUser)

	//name := ctx.PostForm("name")
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	checkdata_bool := utils.CheckData(ctx, name, telephone, password)
	telephone_bool := utils.IsTelephoneExist(telephone)
	if checkdata_bool == false || telephone_bool == false {
		ctx.JSON(422, common.HttpResponse(422, nil, "验证失败"))
		return
	}
	//密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, common.HttpResponse(500, nil, "加密失败"))
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	//创建用户
	common.GetDB().Create(&newUser)

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(500, common.HttpResponse(500, err, "系统出错"))
		log.Printf("toke err:%#v\n", err)
		return
	}

	//返回结果
	ctx.JSON(200, common.HttpResponse(200, token, "token发放成功"))
	ctx.JSON(200, common.HttpResponse(200, newUser, "注册成功"))
}
