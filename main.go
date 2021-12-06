package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
	API "turan.com/WeChat-Private/api"
	"turan.com/WeChat-Private/cache"
	_ "turan.com/WeChat-Private/config"
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

type uvInfo struct {
	Uv        float32   `json:"uv"`
	UvTime    time.Time `json:"uv_time"`
	UvMax     float32   `json:"uv_max"`
	Ozone     float32   `json:"ozone"`
	OzoneTime time.Time `json:"ozone_time"`
}

func main() {

	//setUp()

	//decodeString, _ := base64.StdEncoding.DecodeString(viper.GetString("alibaba.accessKey"))
	//
	API.CorsInfo()

	API.GetCors().GetDucketList()

}
