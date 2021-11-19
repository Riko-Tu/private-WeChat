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
func GetCode() string {
	var codeStr string
	for i := 0; i < 6; i++ {
		code, _ := strconv.Atoi(fmt.Sprintf("%d", rand.Intn(9)))
		codeStr += fmt.Sprintf("%d", code)
	}
	return codeStr
}

var (
	from     = viper.GetString("smtp.from")
	host     = viper.GetString("smtp.qq.host")
	port, _  = strconv.Atoi(viper.GetString("smtp.qq.port"))
	username = viper.GetString("smtp.qq.username")
	password = viper.GetString("smtp.qq.password")
)

//发送邮件
func SendEmail(addressEmail string, code string) error {
	m := gomail.NewMessage()            //获取邮件对象
	m.SetHeader("From", from)           //发件人
	m.SetHeader("To", addressEmail)     //收件人
	m.SetHeader("Subject", "chat【验证码】") //标题
	m.SetBody("text/html", fmt.Sprintf("你的验证码是%s", code))

	d := gomail.Dialer{Host: host, Port: port, Username: username, Password: password}
	return d.DialAndSend(m)
}
