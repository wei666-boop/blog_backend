package utils

import (
	"demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckData(ctx *gin.Context, name string, telephone string, password string) bool {
	if len(telephone) == 11 {
		ctx.JSON(422, common.HttpResponse(422, nil, "手机号只能为11"))
		return false
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, common.HttpResponse(422, nil, "密码少于6位"))
		return false
	}
	if len(name) == 0 {
		name = RandomString(10)
		return true
	}
	return true
}
