package dao

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
