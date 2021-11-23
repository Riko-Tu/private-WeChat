package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	API "turan.com/WeChat-Private/api"
	"turan.com/WeChat-Private/logic"
	"turan.com/WeChat-Private/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//从http请求中获取token
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
		//解析token
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
			ctx.Abort()
			return
		}
		//获取uid存到ctx中
		ctx.Set("uid", claims["uid"])
		ctx.Next()
	}
}

func LogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		host := ctx.Request.URL.Host
		query := ctx.Request.URL.RawQuery
		ctx.Next()
		//等next函数执行完，回来时再存储日志
		ipInfo, err := API.GetIpInfo(host)
		if err != nil {
			zap.L().Debug("", zap.String("path", path),
				zap.String("method", method),
				zap.String("query", query),
				zap.String("host", host),
				zap.String("getIpInfoErr", err.Error()))
			return
		}
		zap.L().Debug("", zap.String("host", host),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("continent", ipInfo.Continent),
			zap.String("country", ipInfo.Country),
			zap.String("regionName", ipInfo.RegionName),
			zap.String("city", ipInfo.City),
			zap.String("org", ipInfo.Org),
			zap.String("isp", ipInfo.Isp),
			zap.Bool("mobile", ipInfo.Mobile),
		)

	}
}
