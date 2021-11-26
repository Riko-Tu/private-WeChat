package route

import (
	"github.com/gin-gonic/gin"
	API "turan.com/WeChat-Private/api"
	"turan.com/WeChat-Private/controller"
)

func route(e *gin.Engine) {
	//用户相关
	user := e.Group("/user")
	user.POST("/emailLogin", controller.EmailLogin)
	user.POST("/getEmailCode", controller.GetEmailCode)
	user.GET("/getUv", LogMiddleware(), controller.GetUv)
	user.POST("/uploadImage", LogMiddleware(), AuthMiddleWare(), controller.UpLoadImage)
	user.POST("/getUser", AuthMiddleWare(), controller.GetUser)
	//其他类别

	a := e.Group("/api", LogMiddleware(), AuthMiddleWare())
	a.GET("/sister", API.GetSister)
}

func SetUp() error {
	r := gin.Default()
	gin.New().Use()
	route(r)
	err := r.Run("127.0.0.1:8080")
	return err
}
