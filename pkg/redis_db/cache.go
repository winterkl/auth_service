package redis_db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	*redis.Client
}

func New(host, password string, dataBase, port int) *Client {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       dataBase,
	}
	redisDB := redis.NewClient(options)
	if err := redisDB.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	return &Client{redisDB}
}
