package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"math/rand"
	"regexp"
	"strconv"
)

//邮箱正则
var validEmail = regexp.MustCompilePOSIX("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")

//邮箱校验
func VerifyEmail(email string) bool {
	return validEmail.MatchString(email)
}

//获取验证码
func getCode() int {
	i := rand.Int()
	return i

}

//发送邮件
func SendEmail(addressEmail string) {
	from := viper.GetString("smtp.from")
	host := viper.GetString("smtp.qq.host")
	port, _ := strconv.Atoi(viper.GetString("smtp.qq.port"))
	username := viper.GetString("smtp.qq.username")
	password := viper.GetString("smtp.qq.password")

	m := gomail.NewMessage()
	m.SetHeader("From", from)           //发件人
	m.SetHeader("To", addressEmail)     //收件人
	m.SetHeader("Subject", "chat【验证码】") //标题
	m.SetBody("text/html", "hello")

	d := gomail.Dialer{Host: host, Port: port, Username: username, Password: password}
	if err := d.DialAndSend(m); err != nil { //发送失败报错
		fmt.Println(err.Error())
	}
}
