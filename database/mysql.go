package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	//gorm需要使用mysql的驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

func SetUp() error {
	var err error
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	host := viper.GetString("private-chat.mysql.host")
	user := viper.GetString("private-chat.mysql.user")
	password := viper.GetString("private-chat.mysql.password")
	dbName := viper.GetString("private-chat.mysql.name")
	charset := viper.GetString("private-chat.mysql.charset")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, dbName, charset)
	db, err = gorm.Open("mysql", dsn)
	zap.L().Debug("Success Connect to dataBase")
	return err

}

func GetDB() *gorm.DB {
	return db
}
