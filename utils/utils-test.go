package utils

import "fmt"

var (
	EmailList = map[string]string{"qq邮箱1": "1234576@qq.com", "qq邮箱2": "7894641@qq.com",
		"gmail邮箱1": "tuasn123@gmail.com", "gmail邮箱2": "qowemmmc@gmail.com",
		"网易邮箱1": "123123@163.com", "网易邮箱2": "124314766@163.com"}
)

//不用类型的邮箱测试
func ReEmailTest(email map[string]string) {
	for k, v := range email {
		fmt.Println(k, VerifyEmail(v))
	}
}
