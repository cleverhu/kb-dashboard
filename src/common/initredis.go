package common

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Rds *redis.Client

func init() {
	fmt.Println("init redis")
	//初始化redis
	Rds = redisClient()
	fmt.Println(Rds)
}

func redisClient() *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     "172.24.212.235:6379",
		Password: "", //
		DB:       0,  //
	})
	//fmt.Println(rds)
	return rds
}