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
		Addr:     "101.132.107.3:6379",
		Password: "", //
		DB:       0,  //
	})
	//fmt.Println(rds)
	return rds
}