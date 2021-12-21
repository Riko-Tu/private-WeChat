package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	API "turan.com/WeChat-Private/api"
	_ "turan.com/WeChat-Private/config"
	"turan.com/WeChat-Private/dao/cache"
	"turan.com/WeChat-Private/dao/database"
	"turan.com/WeChat-Private/log"
	"turan.com/WeChat-Private/route"
)

func TestLevelUnmarshalUnknownText(t *testing.T) {
	var l zapcore.Level
	err := l.UnmarshalText([]byte("debug"))
	if err != nil {
		assert.Contains(t, err.Error(), "unrecognized level", "Expected unmarshaling arbitrary text to fail.")
	}

	fmt.Println(l.String())
}

func setUp() {

	//日志初始化
	err := log.SetUp(viper.GetString("log.mode"))
	if err != nil {
		panic(fmt.Sprintf("logInitErr:%v", err.Error()))
	}
	zap.L().Info("123", zap.String("mna", "123"))

	//数据库初始化
	err = database.SetUp()
	if err != nil {
		panic(fmt.Sprintf("dataBaseInitErr:%v", err.Error()))
	}

	//redis初始化
	cache.SetUp()

	//cors初始化
	API.CorsInfo()

	//路由初始化
	err = route.SetUp()
	if err != nil {
		panic(err.Error())
	}
}

// @title wechat
// @version 1.0
// @description 测试环境
// @termsOfService http://swagger.io/terms/
// @contact.name 徒然
// @contact.url http://www.swagger.io/support
// @contact.email

//@host 127.0.0.1:8080
//@BasePath /

func main() {
	setUp()

}
