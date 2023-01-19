package lib

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
	"twitch_chat_analysis/config"
)

var (
	Connection *rabbitmq.Conn
)

func InitializeRabbitMQConnection() error {
	conn, err := rabbitmq.NewConn(
		config.Get().RabbitMQ.URL,
		rabbitmq.WithConnectionOptionsLogging,
	)

	if err != nil {
		logrus.Error(err)
		return errors.New("could not establish connection with RabbitMQ")
	}

	logrus.Info("rabbitmq connection established!!")
	Connection = conn

	return nil
}
