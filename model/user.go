package model

import (
	"fmt"
	"go.uber.org/zap"
	"time"
	"turan.com/WeChat-Private/database"
	"turan.com/WeChat-Private/utils"
)

type user struct {
	Id         int    `gorm:"primary_key,column:id"`
	Name       string `gorm:"column:name"`
	Uid        string `gorm:"column:uid"`
	Birthday   int64  `gorm:"column:birthday"`
	Area       string `gorm:"column:area"`
	Email      string `gorm:"column:email;default:'22'"`
	Slogan     string `gorm:"column:slogan"`
	Telephone  string `gorm:"column:telephone"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:create_time"` //使用mysql使用int来存储时间，如果使用其他日期类型数据，则该字段不能为空
	UpdateTime int64  `gorm:"column:update_time"`
	DeleteTime int64  `gorm:"column:delete_time"`
}

//返回表名
func getUserTableName() string {
	return "user"
}

func GetId() {
	var id []int
	database.GetDB().Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
	fmt.Println(id[0])
}

//用户注册:插入一条新数据，字段为零值，则不插入该子端
func EmailRegister(email string) error {
	uuid := utils.GetUuid(email)
	var user = user{
		Email:      email,
		Uid:        uuid.String(),
		CreateTime: time.Now().Unix(),
	}
	ups := make(map[string]interface{})
	ups["email"] = email
	ups["uid"] = uuid.String()
	ups["create_time"] = time.Now()

	zap.L().Debug(email)
	err := database.GetDB().Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

//获取用户信息
func GetUserInfo() {

}

//修改手机号

//判断email是否存在
func IsEmailExist(email string) bool {
	var user = &user{}
	//不存在时，返回错误
	err := database.GetDB().Where("email=?", email).First(user).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

//返回密码

//用户修改资料

//更换头像
