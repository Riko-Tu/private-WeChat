package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"turan.com/WeChat-Private/logic"
	"turan.com/WeChat-Private/utils"
)

func EmailLogin(c *gin.Context) {
	c.JSONP(http.StatusOK, 23)
}

func GetEmailCode(c *gin.Context) {
	//通过post表的的key获取参数
	email := c.PostForm("email") //body中添加email字段
	isPass := utils.VerifyEmail(email)
	if isPass {
		//发送邮件
		msg := logic.SendEmail(email)
		logicReply(c, msg)
	}
	CodeMsgReply(c, EmailErr)
}
