package model

import (
	"time"
	"turan.com/WeChat-Private/database"
	"turan.com/WeChat-Private/utils"
)

type user struct {
	id         int 			`gorm:"column:id,type:primaryKey"`
	name       string		`gorm:"column:name"`
	uid        string			`gorm:"column:uid"`
	birthday   time.Time	`gorm:"column:birthday"`
	area       string		`gorm:"column:area"`
	email      string		`gorm:"column:email"`
	slogan     string		`gorm:"column:slogan"`
	telephone  string		`gorm:"column:telephone"`
	password   string		`gorm:"column:password"`
	createTime time.Time	`gorm:"column:create_time"`
	updateTime time.Time	`gorm:"column:update_time"`
	deleteTime time.Time	`gorm:"column:delete_time"`
}

//返回表名
func getUserTableName() string {
	return "user"
}



//用户Email注册
func EmailRegister(email string)  error {
	uuid := utils.GetUuid(email)
	var user =&user{
		email: email,
		uid: uuid.String(),
		createTime: time.Now(),
	}
	err := database.GetDB().Create(user).Error
	return err
}
//获取用户信息

//修改手机号

//判断email是否存在
func IsEmailExist(email string) bool  {
	var user *user
	database.GetDB().Where("email=?",email).First(getUserTableName()).Scan(user)
	if user.id <=0 {
		return false
	}
	return true
}


//返回密码

//用户修改资料

//更换头像
