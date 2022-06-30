package cache

import (
	"fmt"
	"ginApi/conf"
	"github.com/go-redis/redis"
	"time"
)

var RedisClient *redis.Client

func InitRedis(config *conf.RedisConfig)  {
	addr := config.Host+ ":" +config.Port
	RedisClient = redis.NewClient(&redis.Options{
		Addr:    addr, // Redis地址
		Password: config.Password,  // Redis账号
		DB:       config.Db,   // Redis库
		PoolSize: 10,  // Redis连接池大小
		MaxRetries: 3,              // 最大重试次数
		IdleTimeout: 10*time.Second,            // 空闲链接超时时间
	})
	pong, err := RedisClient.Ping().Result()
	if err == redis.Nil {
		fmt.Println("Redis异常")
	} else if err != nil {
		fmt.Println("失败:",err)
	} else {
		fmt.Println(pong)
	}
}