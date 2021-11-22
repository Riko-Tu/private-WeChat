package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TokenExpiresTime = 60 * 60 * 24 * 7 //token过期时间为一周

type claims struct {
	uid   string
	email string
	jwt.StandardClaims
}

//创建token
func CreateToken(uid string) {
	claims := &claims{
		uid:   uid,
		email: "",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpiresTime * time.Second).Unix(),
		},
	}
	/**
	token构成：
			 header: 存放token的类型  type：jwt; alg：加密方式
			 claims: 主要放置用户信息，和其他发行人等需要的字段
			 method: 密钥
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//token包含了，header与claims，通过singnedString进行密钥加密
	signedString, err := token.SignedString([]byte("1231-23-21"))
	if err != nil {
		panic(err.Error())
	}
	//token, err := jwt.ParseWithClaims(token)
	fmt.Println(signedString)
}

//解析token
func ParseToken() {

}

//更新token过期时间
func UpdateTokenExpiresTime() {

}
