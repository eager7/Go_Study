package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

func InitRedis(addr, pass string, zone int) *redis.Client {
	fmt.Println("init redis:", addr, pass, zone)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       zone,
	})
}
