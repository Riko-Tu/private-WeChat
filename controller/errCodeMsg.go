package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type errCodeMsg struct {
	code int
	msg  string
}

func (e *errCodeMsg) toString() {
	fmt.Println(fmt.Sprintf("%d %v", e.code, e.msg))
}

var (
	//邮箱验证错误
	emailErr = errCodeMsg{code: 1001, msg: "email error"}
)

func errMsg(ctx *gin.Context, err errCodeMsg) {
	ctx.JSON(http.StatusOK, err)
}
