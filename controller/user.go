package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	API "turan.com/WeChat-Private/api"
	"turan.com/WeChat-Private/dao/cache"
	"turan.com/WeChat-Private/logic"
	"turan.com/WeChat-Private/model"
	"turan.com/WeChat-Private/utils"
)

func EmailLogin(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	//邮箱校验
	isPass := utils.VerifyEmail(email)
	getcode, err := cache.GetCode(email)
	if err != nil {
		CodeMsgReply(c, CodeErr)
		return
	}
	//如果邮箱和验证码都会真
	if isPass || code == getcode {
		//为真，进行登录
		msg := logic.EmailLogin(email)
		logicReply(c, msg)
		return
	}
	//校验错误返回，邮箱格式错误
	CodeMsgReply(c, EmailErr)
}

func GetEmailCode(c *gin.Context) {
	//通过post表的的key获取参数
	email := c.PostForm("email") //body中添加email字段
	isPass := utils.VerifyEmail(email)
	if isPass {
		//发送邮件
		msg := logic.SendEmail(email)
		logicReply(c, msg)
		return
	}
	//校验错误返回，邮箱格式错误
	CodeMsgReply(c, EmailErr)
}

//
func GetUv(ctx *gin.Context) {
	lat, _ := ctx.Get("lat")
	lon, _ := ctx.Get("lon")
	uv, err := API.GetUv(lat.(float32), lon.(float32))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"uv": uv[0],
		"uv_time":     uv[1],
		"uv_max":      uv[2],
		"uv_max_time": uv[3]})
}

//上传头像
func UpLoadImage(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	zap.L().Debug(file.Filename)
	//dst：upload文件夹必须存在
	fileName := uid.(string)[:6] + strings.Replace(file.Filename, " ", "", -1)
	dst := "upload/" + fileName

	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	//根据uid存储数据库
	fileUrl:=fmt.Sprintf("http://127.0.0.1:8080/user/getImage/%s", fileName)
	err = model.ImageUpload(fileUrl, uid.(string))
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"msg":"上传成功","Url": fileUrl} )
}

//获取用户头像
func GetUserImage(ctx *gin.Context) {
	//uid, _ := ctx.Get("uid")
	filePath := ctx.Param("path")
	ctx.File("./upload/"+filePath)

}

//获取用户信息
func GetUser(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	user, err := model.GetUserById(uid.(string))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
