package cache

import (
	"time"
	"turan.com/WeChat-Private/cache"
)

//存验证码到redis
func SaveCode(email string, code string) error {
	err := cache.GetRdb().Set(email, code, 1*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

//通过email获取验证码
func GetCode(email string) (string, error) {
	result, err := cache.GetRdb().Get(email).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

//token存储
func SaveTokenByUid(uid string, token string) error {
	err := cache.GetRdb().Set(uid, token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
