package logic

type LogicMsg struct {
	Code int
	Msg  string
}

var (
	//发送邮件失败
	SendEmailFailed = LogicMsg{3001, "send email failed"}
	//code存储失败
	CodeSaveFailed = LogicMsg{3002, "save failed"}
	//创建token失败
	CreateTokenFailed = LogicMsg{3003, "CreateTokenFailed"}
	//存储token失败
	SaveTokenFailed = LogicMsg{3004, "SaveTokenFailed"}
	//token解析失败
	TokenParseFailed = LogicMsg{3005, ""}
)
var (
	//发送邮件成功
	SendEmailSuccess = LogicMsg{4001, "Send Email Success"}
	//返回token
	CreateTookenSuccess = LogicMsg{4002, ""}
)
