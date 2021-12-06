package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

//初始化redis
func SetUp() {
	options := &redis.Options{
		Addr: viper.GetString("redis.host"),
		DB:   0, //数据库号码
	}
	rdb = redis.NewClient(options)
}

//获取客户端
func GetRdb() *redis.Client {
	return rdb
}
