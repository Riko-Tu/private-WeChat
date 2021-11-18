package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func SetUp() {
	options := &redis.Options{
		Addr: viper.GetString("redis.host"),
		DB:   0, //数据库号码
	}
	rdb := redis.NewClient(options)
	rdb.Set()
}
