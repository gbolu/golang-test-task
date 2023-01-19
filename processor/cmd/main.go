package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"twitch_chat_analysis/config"
	"twitch_chat_analysis/lib"
)

const (
	DEFAULT_KEY_STORE = "DEFAULT_KEY_STORE"
)

func main() {
	config.InitializeEnvironmentVariables()
	err := lib.InitializeRabbitMQConnection()
	if err != nil {
		logrus.Info(err)
		return
	}

	redis := lib.NewRedis()
	arr, _ := redis.Get(DEFAULT_KEY_STORE)
	if arr == "" {

	}

	consumer, err := rabbitmq.NewConsumer(
		lib.Connection,
		func(d rabbitmq.Delivery) rabbitmq.Action {
			logrus.Info(fmt.Sprintf("consumed: %s", string(d.Body)))

			go redis.Set(DEFAULT_KEY_STORE, d.Body)
			return rabbitmq.Ack
		},
		config.MESSAGE_QUEUE_NAME,
		rabbitmq.WithConsumerOptionsRoutingKey(config.MESSAGE_QUEUE_NAME),
		rabbitmq.WithConsumerOptionsExchangeName(config.MESSAGE_QUEUE_NAME),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("stopping consumer")
}
