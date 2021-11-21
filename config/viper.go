package config

import "github.com/spf13/viper"

func init() {
	/**
	viper使用：
			1.指定配置文件路径
			2.指定配置文件名称
			3.指定配置文件类型
			4.读取配置文件，读取后可以viper在全局使用getString
	*/
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

}
