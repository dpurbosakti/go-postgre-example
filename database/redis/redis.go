package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RdClient *redis.Client

func RegisterRedis() {
	//Initializing redis
	// dsn := viper.GetString("redisConf.dsn")
	// if len(dsn) == 0 {
	dsn := "127.0.0.1:6379"
	// }

	RdClient = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := RdClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
