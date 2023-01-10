package config

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func CacheConnection(cacheUrl string, password string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:        cacheUrl,
		Password:    password,
		DB:          0,
		DialTimeout: time.Second * 5,
	})

	// create context
	ctx := context.Background()

	// ping
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("redis connection error: ", err)
	}

	return rdb

}
