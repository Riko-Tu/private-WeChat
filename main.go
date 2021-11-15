package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
	"turan.com/WeChat-Private/middleware"
)

func TestLevelUnmarshalUnknownText(t *testing.T) {
	var l zapcore.Level
	err := l.UnmarshalText([]byte("debug"))
	if err != nil {
		assert.Contains(t, err.Error(), "unrecognized level", "Expected unmarshaling arbitrary text to fail.")
	}

	fmt.Println(l.String())
}

func init() {
	//日志初始化
	err := middleware.LoggerInit("staging")
	if err != nil {
		panic(fmt.Sprintf("logInitErr:%v", err.Error()))
	}
	zap.L().Info("123", zap.String("mna", "123"))
}

func main() {
	TestLevelUnmarshalUnknownText(&testing.T{})

}
