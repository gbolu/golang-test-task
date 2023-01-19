package lib

import (
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type (
	IPublisher interface {
		Publish(data []byte) error
	}

	Publisher struct {
		RabbitMQPublisher *rabbitmq.Publisher
		QueueName         string
	}
)

func createRabbitMQPublisher(queueName string) (*rabbitmq.Publisher, error) {
	publisher, err := rabbitmq.NewPublisher(Connection,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(queueName),
		rabbitmq.WithPublisherOptionsExchangeDeclare)

	if err != nil {
		return nil, err
	}

	return publisher, nil
}

func NewPublisher(queueName string) IPublisher {
	publisher, err := createRabbitMQPublisher(queueName)
	if err != nil {
		logrus.Error("error occurred while creating publisher: ", err.Error())
		return &Publisher{}
	}

	return &Publisher{
		RabbitMQPublisher: publisher,
		QueueName:         queueName,
	}
}

func (p Publisher) Publish(data []byte) error {
	err := p.RabbitMQPublisher.Publish(data, []string{p.QueueName},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange(p.QueueName))
	if err != nil {
		return err
	}

	return nil
}
