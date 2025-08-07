package middleware

import (
	"demo/common"
	"demo/model"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" || strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(401, common.HttpResponse(401, nil, "权限不足"))
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		//验证解析出的token是否有效
		token, claims, err := common.ParserToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(401, common.HttpResponse(401, nil, "权限不足"))
			ctx.Abort()
			return
		}

		//验证通过后获取的userid
		userID := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.Where("id=?", userID).First(&user)
		//验证用户是否存在
		if user.ID == 0 {
			ctx.JSON(422, common.HttpResponse(422, nil, "用户不存在"))
			ctx.Abort()
			return
		}
		//将用户信息存放进上下文方便日后读取
		ctx.Set("user", user)
		ctx.Next()
	}
}
