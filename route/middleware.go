package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"turan.com/WeChat-Private/logic"
	"turan.com/WeChat-Private/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		zap.L().Debug(tokenString)
		//检测token是否为空  ||  是否以bearer 开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusOK, gin.H{"code": logic.TokenParseFailed.Code, "msg": logic.TokenParseFailed.Msg})
			//token错误中断请求
			ctx.Abort()
			return
		}
		//去除bearer
		tokenString = tokenString[len("bearer "):]
		zap.L().Debug(tokenString)
		token, err := utils.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
			ctx.Abort()
			return
		}
		//获取uid存到ctx中
		ctx.Set("uid", token["uid"])
		ctx.Next()
	}
}
