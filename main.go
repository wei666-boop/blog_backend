package main

import (
	"demo/common"
	"demo/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func init() {
	workDir, _ := os.Getwd() //读取工作路径
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/")
	err := viper.ReadInConfig() //读取配置文件
	if err != nil {
		panic(err)
	}
}
func main() {
	common.InitDB()
	r := gin.Default()
	r = router.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	}
}
