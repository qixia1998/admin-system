// redis初始化连接

package redis

import (
	"admin-api-golang/common/config"
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	RedisDb *redis.Client
)

// SetupRedisDb Initialize the Redis instance
func SetupRedisDb() error {
	var ctx = context.Background()
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Address,
		Password: config.Config.Redis.Password,
		DB:       0,
	})
	_, err := RedisDb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}
