package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client
var userConnRdb *redis.Client

//初始化redis
func SetUp() {
	options := &redis.Options{
		Addr: viper.GetString("redis.host"),
		DB:   0, //数据库号码
	}
	rdb = redis.NewClient(options)
	option := &redis.Options{
		Addr: viper.GetString("redis.host"),
		DB:   1, //数据库号码
	}
	userConnRdb = redis.NewClient(option)
}

//获取客户端
func GetRdb() *redis.Client {
	return rdb
}
func GetConnDB() *redis.Client {
	return userConnRdb
}
