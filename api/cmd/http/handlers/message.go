package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"twitch_chat_analysis/config"
	"twitch_chat_analysis/lib"
)

type (
	MessageHandler struct {
		Publisher lib.IPublisher
	}

	IMessageHandler interface {
		ProcessMessage(c *gin.Context)
	}
)

func NewMessageHandler() IMessageHandler {
	return &MessageHandler{
		Publisher: lib.NewPublisher(config.MESSAGE_QUEUE_NAME),
	}
}

func (m *MessageHandler) ProcessMessage(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Error(err)
		c.JSON(400, "invalid data")
		return
	}

	publishErr := m.Publisher.Publish(jsonData)
	if publishErr != nil {
		logrus.Error("error occurred while publishing message: ", publishErr)
		c.JSON(500, "could not process message. please try again")
		return
	}

	c.JSON(200, "success")
}
