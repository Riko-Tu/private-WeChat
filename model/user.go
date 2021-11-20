package model

import (
	"time"
	"turan.com/WeChat-Private/database"
)

type user struct {
	id         int 			`gorm:"column:id,type:primaryKey"`
	name       string		`gorm:"column:name"`
	uid        int8			`gorm:"column:uid"`
	birthday   time.Time	`gorm:"column:birthday"`
	area       string		`gorm:"column:area"`
	email      string		`gorm:"column:email"`
	slogan     string		`gorm:"column:slogan"`
	telephone  string		`gorm:"column:telephone"`
	password   string		`gorm:"column:"`
	createTime time.Time	`gorm:"column:"`
	updateTime time.Time	`gorm:"column:"`
	deleteTime time.Time	`gorm:"column:"`
}

//返回表名
func getUserTableName() string {
	return "user"
}

//用户Email注册
func EmailRegister()  {
	database.GetDB().Model(getUserTableName()).

}
//获取用户信息

//修改手机号

//判断用户是否存在

//返回密码

//用户修改资料

//更换头像
