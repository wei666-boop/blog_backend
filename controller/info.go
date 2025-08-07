package controller

import (
	"demo/common"
	"demo/model"
	"github.com/gin-gonic/gin"
)

// 获得用户信息
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, common.HttpResponse(200, user.(model.User), "用户信息"))
}

//由于通过上下文传递的数据为any类型所以要先通过断言转化为我们想要的类型
