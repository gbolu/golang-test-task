package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

type (
	EnvType struct {
		RabbitMQ RabbitMQEnv
		DATABASE DatabaseEnv
	}

	RabbitMQEnv struct {
		URL string `json:"-"`
	}

	RedisEnv struct {
		REDIS_URL string `json:"-"`
	}

	DatabaseEnv struct {
		REDIS RedisEnv
	}
)

var (
	env EnvType
)

const (
	MESSAGE_QUEUE_NAME = "messages"
)

func getEnv(Name, Default string) string {
	if v := os.Getenv(Name); v != "" {
		return v
	}

	return Default
}

func mustGetEnv(Name string) string {
	if v := os.Getenv(Name); v != "" {
		return v
	}

	panic(fmt.Sprintf("%s not found in env", Name))
}

func Get() EnvType {
	_env := env
	return _env
}

func InitializeEnvironmentVariables() {
	if er := godotenv.Load(".env"); er != nil {
		logrus.Info(er.Error())
	}

	env.RabbitMQ.URL = mustGetEnv("RABBITMQ_URL")
	env.DATABASE.REDIS.REDIS_URL = mustGetEnv("REDIS_URL")
}
