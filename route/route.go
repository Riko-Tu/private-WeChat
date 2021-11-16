package route

import (
	"github.com/gin-gonic/gin"
	"turan.com/WeChat-Private/controller"
)

func route(e *gin.Engine) {
	user := e.Group("/user")
	user.POST("/emailLogin", controller.EmailLogin)
	user.GET("/getEmailCode/:email", controller.GetEmailCode)
	user.GET("/g")
}

func SetUp() error {
	r := gin.Default()
	route(r)
	err := r.Run("127.0.0.1:8080")
	return err
}
