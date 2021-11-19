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
)
var (
	//发送邮件成功
	SendEmailSuccess = LogicMsg{4002, "Send Email Success"}
)
