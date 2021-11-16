package log

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strconv"
	"time"
)

/**
日志使用：
		1.设置Encoder配置
		2.encoder传入zap.core
		3.将实现了io.write接口的writer传入zap.core
		4.配置zap.core的日志等级
		5.使用zap.core获取logger对象，并注册全局zap
*/

//日志初始化对象
func SetUp(mode string) error {
	//case "debug", "DEBUG":
	//	case "info", "INFO", "": // make the zero value useful
	//	case "warn", "WARN":
	//	case "error", "ERROR":
	//	case "dpanic", "DPANIC":
	//	case "panic", "PANIC":
	//	case "fatal", "FATAL":
	var l zapcore.Level
	//反射text进行配置日志等级
	err := l.UnmarshalText([]byte(viper.GetString("log.level")))
	if err != nil {
		return err
	}
	coreConfig := getZapCore(mode, getJsonEncoder(), l)
	logger := zap.New(coreConfig, zap.AddCaller())
	//注冊全局log可通过zap.l()获取
	zap.ReplaceGlobals(logger)

	return nil
}

//分割日志配置
func getWriter() zapcore.WriteSyncer {
	fileName := fmt.Sprintf("./log/log%s.txt", time.Now().String()[:11])
	MaxSize, err := strconv.Atoi(viper.GetString("log.MaxSize"))
	if err != nil {
		panic("ConfigErr:log.MaxSize NOT integer")
	}
	MaxAge, err := strconv.Atoi(viper.GetString("log.MaxAge"))
	if err != nil {
		panic("ConfigErr:log.MaxAge NOT integer")
	}

	writer := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    MaxSize,
		MaxBackups: 0,
		MaxAge:     MaxAge,
		LocalTime:  true,
		Compress:   false,
	}
	return zapcore.AddSync(writer)
}

/**
获取zap.core配置

*/
func getZapCore(mode string, config zapcore.Encoder, logLevel zapcore.LevelEnabler) zapcore.Core {
	if mode == "dev" {
		zapCore := zapcore.NewCore(config, zapcore.AddSync(os.Stdout), logLevel)
		return zapCore
	} else if mode == "staging" {
		zapCore := zapcore.NewCore(config, zapcore.NewMultiWriteSyncer(getWriter(), zapcore.AddSync(os.Stdout)), logLevel)
		return zapCore
	} else {
		zapCore := zapcore.NewCore(config, getWriter(), logLevel)
		return zapCore
	}
}

//实现timeEncoder类型的函数，格式化时间
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

//获取Encoder配置
func getJsonEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder

	//newEncoderConfig := zapcore.EncoderConfig{
	//	MessageKey:       "msg",    //debug("asd"）,info("123",...)   asd与123默认的key为msg,必须不写不输出
	//	LevelKey:         "level",  //日志等级的key  ,不写不输出
	//	TimeKey:          "date",   //时间key命名
	//	NameKey:          "debug",  //日志的命名
	//	CallerKey:        "caller", //这条日志的输入的函数key,不写不输出，必须写，便于找位子
	//	FunctionKey:      "",       //表示主入口的key  main.main,一定不需要
	//	StacktraceKey:    "S",
	//	LineEnding:       zapcore.DefaultLineEnding,   //行以"/n"结束
	//	EncodeLevel:      zapcore.CapitalLevelEncoder, //序列化输出并修改颜色 ,小写输出
	//	EncodeTime:       timeEncoder,                 //设置时间的输出格式
	//	EncodeDuration:   zapcore.StringDurationEncoder,
	//	EncodeCaller:     zapcore.ShortCallerEncoder, //调用函数的函数命名格式
	//	ConsoleSeparator: "-",
	//}
	return zapcore.NewJSONEncoder(encoderConfig)

}
