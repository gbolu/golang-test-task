package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"twitch_chat_analysis/api/cmd/http/handlers"
	"twitch_chat_analysis/config"
	"twitch_chat_analysis/lib"
)

func main() {
	config.InitializeEnvironmentVariables()
	connectionErr := lib.InitializeRabbitMQConnection()
	if connectionErr != nil {
		logrus.Error(connectionErr)
		return
	}

	r := gin.Default()
	messageHandler := handlers.NewMessageHandler()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	r.POST("/message", messageHandler.ProcessMessage)

	err := r.Run()
	if err != nil {
		logrus.Error()
		return
	}
}
