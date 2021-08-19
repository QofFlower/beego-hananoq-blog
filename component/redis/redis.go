package redis

import (
	"beego-hananoq-blog/conf"
	"github.com/astaxie/beego/logs"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	addr := conf.GetConfig().Redis.Default.Addr
	db := conf.GetConfig().Redis.Default.Db
	redisClient = redis.NewClient(
		&redis.Options{
			Addr: addr,
			DB:   db,
		},
	)
	result, err := redisClient.Ping().Result()
	if err != nil {
		logs.Error("Redis init error: ", err)
		return
	}
	logs.Debug("Result for redis init: ", result)
}

func GetRedisClient() *redis.Client {
	return redisClient
}
