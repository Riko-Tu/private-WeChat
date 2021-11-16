package utils

import (
	"regexp"
)

//邮箱正则
var validEmail = regexp.MustCompilePOSIX("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")

//邮箱校验
func VerifyEmail(email string) bool {
	return validEmail.MatchString(email)
}
