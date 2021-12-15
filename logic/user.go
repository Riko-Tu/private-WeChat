package logic

import (
	"turan.com/WeChat-Private/dao/cache"
	"turan.com/WeChat-Private/model"
	"turan.com/WeChat-Private/utils"
)

//发送邮件
func SendEmail(email string) LogicMsg {
	//获取验证码
	code := utils.GetCode()
	//校验邮箱是否存在
	isExist := model.IsEmailExist(email)
	if !isExist {

		//存储email与uid
		err := model.EmailRegister(email)
		if err != nil {
			SendEmailFailed.Msg = err.Error()
			return SendEmailFailed
		}

	}

	//存储验证码
	err := cache.GetRdb().SaveCode(email, code)
	if err != nil {
		return CodeSaveFailed
	}
	//发送邮件
	err = utils.SendEmail(email, code)
	if err != nil {
		SendEmailFailed.Msg = err.Error()
		return SendEmailFailed
	}
	return SendEmailSuccess
}

//邮件登录
func EmailLogin(email string) (error LogicMsg) {
	//通过email查询uid
	user := model.GetUid(email)
	//创建token
	token, err := utils.CreateToken(user.Uid)
	//创建token错误
	if err != nil {
		return CreateTokenFailed
	}
	//通过uid存储token
	err = cache.GetRdb().SaveTokenByUid(user.Uid, token)
	//存储失败
	if err != nil {
		return SaveTokenFailed
	}
	CreateTookenSuccess.Msg = token
	return CreateTookenSuccess
}
