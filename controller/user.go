package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"turan.com/WeChat-Private/utils"
)

func EmailLogin(c *gin.Context) {
	c.JSONP(http.StatusOK, 23)
}

func GetEmailCode(c *gin.Context) {
	//通过post表的的key获取参数
	email := c.PostForm("email")
	isPass := utils.VerifyEmail(email)
	if isPass == false {
		errMsg(c, emailErr)
	}
	c.JSONP(http.StatusOK, gin.H{"email": email})
}
