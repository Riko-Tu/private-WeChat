package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

var rdb *redisDb
var userConnRdb *redisDb

type redisDb struct {
	HostAndPort string
	DbNumber    int
	Auth        string
	client      *redis.Client
}

func NewRdbClient(db *redisDb) {
	options := &redis.Options{Addr: db.HostAndPort, DB: db.DbNumber}
	db.client = redis.NewClient(options)
	rdb = db
	options.DB = 1
	db.client = redis.NewClient(options)
	userConnRdb = db

}

//初始化redis
func SetUp() {

	redisClient := &redisDb{
		HostAndPort: viper.GetString("redis.host"),
		DbNumber:    0,
		Auth:        "",
	}
	NewRdbClient(redisClient)

}

//获取客户端
func GetRdb() *redisDb {
	return rdb
}

//连接池
func GetConnDB() *redisDb {
	return userConnRdb
}

//存验证码到redis
func (r *redisDb) SaveCode(email string, code string) error {
	err := r.client.Set(email, code, 1*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

//通过email获取验证码
func (r *redisDb) GetCode(email string) (string, error) {
	result, err := r.client.Get(email).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

//token存储
func (r *redisDb) SaveTokenByUid(uid string, token string) error {
	err := r.client.Set(uid, token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
