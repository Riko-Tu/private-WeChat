package controller

import (
	"github.com/gin-gonic/gin"
	"turan.com/WeChat-Private/dao/cache"
	"turan.com/WeChat-Private/logic"
	"turan.com/WeChat-Private/utils"
)

func EmailLogin(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	//邮箱校验
	isPass := utils.VerifyEmail(email)
	getcode, err := cache.GetCode(email)
	if err != nil {
		CodeMsgReply(c, CodeErr)
		return
	}
	//如果邮箱和验证码都会真
	if isPass || code == getcode {
		//为真，进行登录
		msg := logic.EmailLogin(email)
		logicReply(c, msg)
		return
	}
	//校验错误返回，邮箱格式错误
	CodeMsgReply(c, EmailErr)
}

func GetEmailCode(c *gin.Context) {
	//通过post表的的key获取参数
	email := c.PostForm("email") //body中添加email字段
	isPass := utils.VerifyEmail(email)
	if isPass {
		//发送邮件
		msg := logic.SendEmail(email)
		logicReply(c, msg)
		return
	}
	//校验错误返回，邮箱格式错误
	CodeMsgReply(c, EmailErr)
}
