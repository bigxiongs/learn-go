package main

import (
	"github.com/gin-gonic/gin"
	"learn/bootstrap"
	"learn/global"
)

func main() {
	bootstrap.InitializeConfig()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.Run(":" + global.App.Config.App.Port)
}
