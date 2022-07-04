package cache

import "github.com/go-redis/redis/v8"

func GetClient() *redis.Client {
	return cli
}
