package logic

import (
	"turan.com/WeChat-Private/dao"
	"turan.com/WeChat-Private/utils"
)

func SendEmail(email string) LogicMsg {
	//获取验证码
	code := utils.GetCode()
	//存储验证码
	err := dao.SaveCode(email, code)
	if err != nil {
		return CodeSaveFailed
	}
	//发送邮件
	err = utils.SendEmail(email, code)
	if err != nil {
		return SendEmailFailed
	}
	return SendEmailSuccess
}
