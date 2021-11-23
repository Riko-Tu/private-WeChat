package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"turan.com/WeChat-Private/cache"
	_ "turan.com/WeChat-Private/config"
	"turan.com/WeChat-Private/dao/database"
	"turan.com/WeChat-Private/log"
	"turan.com/WeChat-Private/route"
	"turan.com/WeChat-Private/utils"
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

	//路由初始化
	err = route.SetUp()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	//setUp()
	//
	//token, err := utils.CreateToken("7c67e146-a923-3f42-9d8e-e7ab8a3fea5e")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//
	//fmt.Println(token[:])
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mzc2NDI2NTUsInVpZCI6IjdjNjdlMTQ2LWE5MjMtM2Y0Mi05ZDhlLWU3YWI4YTNmZWE1ZSJ9.vDZBC64lytfes_U_XHfGbSkoVdgEqjl5EbitmFrYhj0"
	err := utils.ParseToken(token)
	if err != nil {
		fmt.Println(err.Error())
	}
}
