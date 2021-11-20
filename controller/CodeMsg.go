package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"turan.com/WeChat-Private/logic"
)

type CodeMsg struct {
	Code int
	Msg  string
}

func (e *CodeMsg) toString() {
	fmt.Println(fmt.Sprintf("%d %v", e.Code, e.Msg))
}

//校验类
var (
	//邮箱验证错误
	EmailErr = CodeMsg{1001, "email error"}
)

//codeMsgReply
func CodeMsgReply(ctx *gin.Context, msg CodeMsg) {
	ctx.JSON(http.StatusOK, gin.H{"code": msg.Code, "msg": msg.Msg})
}

//codeMsgReply
func logicReply(ctx *gin.Context, msg logic.LogicMsg) {
	ctx.JSON(http.StatusOK, gin.H{"code": msg.Code, "msg": msg.Msg})
}
