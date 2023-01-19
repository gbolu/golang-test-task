package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"twitch_chat_analysis/config"
)

func main() {
	config.InitializeEnvironmentVariables()
	r := gin.Default()
	r.GET("/message/list", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	err := r.Run()
	if err != nil {
		logrus.Error()
		return
	}
}
