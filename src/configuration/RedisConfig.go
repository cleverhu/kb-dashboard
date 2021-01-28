package configuration

import (
	"github.com/go-redis/redis"
	"knowledgeBase/src/common"
)

type RedisConfig struct {
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{}
}

func (this *RedisConfig) RedisClient() *redis.Client {

	return common.Rds
}
