package utils

import (
	"fmt"
	uuid2 "github.com/satori/go.uuid"
)

var (
	//邮箱测试列表
	EmailList = map[string]string{"qq邮箱1": "1234576@qq.com", "qq邮箱2": "7894641@qq.com",
		"gmail邮箱1": "tuasn123@gmail.com", "gmail邮箱2": "qowemmmc@gmail.com",
		"网易邮箱1": "123123@163.com", "网易邮箱2": "124314766@163.com"}

	//uuid测试列表
	uuidList = []string{"1231", "!@#@!$#", "123123@163.com", "124314766@163.com", "arafwr@#$"}
)

//不用类型的邮箱测试
func ReEmailTest() {
	for k, v := range EmailList {
		fmt.Println(k, VerifyEmail(v))
	}
}

//测试不同输入类型数据
func GetUuidTest() {
	var uuid uuid2.UUID
	for i := 0; i < len(uuidList); i++ {
		uuid = GetUuid(uuidList[i])
		fmt.Println(uuid.String())
	}

}
