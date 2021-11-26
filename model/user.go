package model

import (
	"fmt"
	"time"
	"turan.com/WeChat-Private/dao/database"
	"turan.com/WeChat-Private/utils"
)

type user struct {
	Id         int    `gorm:"primary_key,column:id"`
	Name       string `gorm:"column:name"`
	Uid        string `gorm:"column:uid"`
	Image      string `gorm:"column:image"`
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

func GetUserStruct() *user {
	return &user{}
}

//返回表名
func (*user) TableName() string {
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
	//只插入邮箱，uid，和创建时间
	var user = user{
		Email:      email,
		Uid:        uuid.String(),
		CreateTime: time.Now().Unix(),
	}

	err := database.GetDB().Create(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

//上传图片
func ImageUpload(imagePath, uid string) error {
	//结构体内部可以添加where条件

	user := &user{Uid: uid}
	err := database.GetDB().Model(user).Update("image", imagePath).Error
	if err != nil {
		return err
	}
	return nil
}

//获取用户信息
func GetUserById(uid string) (*user, error) {
	var user = &user{}
	err := database.GetDB().Where("uid = ?", uid).First(user).Scan(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

//修改手机号

//判断email是否存在
func IsEmailExist(email string) bool {
	var user = &user{}
	//不存在时，返回错误
	err := database.GetDB().Where("email=?", email).First(user).Error
	if err != nil {
		return false
	}
	return true
}

//通过email获取uid
func GetUid(email string) *user {
	var user = &user{}
	database.GetDB().Where("email = ?", email).First(user).Scan(user)
	return user
}

//返回密码

//用户修改资料

//更换头像
