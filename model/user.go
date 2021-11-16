package model

import "time"

type user struct {
	id         int
	name       string
	uid        int8
	birthday   time.Time
	area       string
	email      string
	slogan     string
	telephone  string
	password   string
	createTime time.Time
	updateTime time.Time
	deleteTime time.Time
}

//生成UID

//用户email注册

//获取用户信息

//修改手机号

//判断用户是否存在

//返回密码

//用户修改资料

//更换头像
