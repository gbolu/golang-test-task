package lib

import (
	"context"
	"crypto/tls"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"strings"
	"twitch_chat_analysis/config"
)

var GlobalRedisClient *redis.Client = nil

type Redis struct {
	Client *redis.Client
}

func (r *Redis) Set(key string, value interface{}) error {
	err := r.Client.Set(context.Background(), key, value, redis.KeepTTL).Err()
	if err != nil { //
		return err
	}

	logrus.Info("stored value in ", key)
	return nil
}

func (r *Redis) Get(key string) (string, error) {
	value, err := r.Client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}

func (r *Redis) Del(key string) (int64, error) {
	value, err := r.Client.Del(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func GetRedis() *Redis {
	if GlobalRedisClient != nil {
		r := Redis{Client: GlobalRedisClient}
		return &r
	}
	//Initializing redis
	redisAddress := config.Get().DATABASE.REDIS.REDIS_URL

	redisOptions, err := redis.ParseURL(redisAddress)

	if err != nil {
		panic(err)
	}

	if strings.Contains(redisAddress, "rediss") {
		redisOptions.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client := redis.NewClient(redisOptions)
	result, err := client.Ping(context.Background()).Result()

	println(result)

	if err != nil {
		panic(err)
	}

	GlobalRedisClient = client

	r := Redis{Client: GlobalRedisClient}
	return &r
}

func NewRedis() *Redis {
	//Initializing redis
	redisAddress := config.Get().DATABASE.REDIS.REDIS_URL

	redisOptions, err := redis.ParseURL(redisAddress)

	if err != nil {
		panic(err)
	}

	if strings.Contains(redisAddress, "rediss") {
		redisOptions.TLSConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client := redis.NewClient(redisOptions)
	result, err := client.Ping(context.Background()).Result()

	println(result)

	if err != nil {
		panic(err)
	}

	r := Redis{Client: client}
	return &r
}
