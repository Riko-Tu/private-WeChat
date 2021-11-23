package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

const (
	TokenExpiresTime = 60 * 60 * 24 * 7 //token过期时间为一周

)

var (
	secret = viper.GetString("secret")
)

//
type claims struct {
	uid   string
	email string
	jwt.StandardClaims
}

//创建token
func CreateToken(uid string) (string, error) {

	//使用了mapClaims类型创建的，该claims可自定义mapKey较为方便
	//还有一个StandardClaims类型是规定了字段的结构体
	claims := &jwt.MapClaims{
		"uid": uid,

		//该字段是过期时间
		"exp": time.Now().Add(TokenExpiresTime * time.Second).Unix(),
	}
	/**
	token构成：
			 header: 存放token的类型  type：jwt; alg：加密方式
			 claims: 主要放置用户信息，和其他发行人等需要的字段
			 method: 密钥
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//token包含了，header与claims，通过singnedString进行密钥加密
	TokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

//解析token
func ParseToken(tokenString string) (claims jwt.Claims, err error) {
	//解析token时，使用mapClaims类型来解析token的Claims类型
	token, err := jwt.ParseWithClaims(tokenString, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		//token使用byte加密，这里使用byte解密
		return []byte(secret), nil
	})
	//解析失败返回错误
	if err != nil {
		zap.L().Debug(err.Error())
		return nil, err
	}
	//将解析好的claims类型转成mapClaims，因为mapClaims是claims的实现类
	mapClaims := token.Claims.(jwt.MapClaims)
	return mapClaims, nil
}

//获取token

//更新token过期时间
func UpdateTokenExpiresTime() {

}
