package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
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
	getcode, err := cache.GetRdb().GetCode(email)
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

	split := strings.Split(file.Filename, ".")
	subffix := split[len(split)-1]
	//获取ocrs
	log.Println(subffix, file.Size)
	bytes := make([]byte, 1024000000)
	cors := API.GetCors()
	open, err := file.Open()
	number, err := open.Read(bytes)
	if err != nil {
		ctx.String(http.StatusOK, "文件读取错误："+err.Error())
		return
	}
	fileName := uid.(string)[:6] + strings.Replace(file.Filename, " ", "", -1)
	if strings.EqualFold(strings.ToLower(subffix), strings.ToLower("png")) ||
		strings.EqualFold(strings.ToLower(subffix), strings.ToLower("png")) {
		//存入图库
		corsFilePath := viper.GetString("alibaba.cors.chatFileDir") + fileName
		err := cors.UploadFile(corsFilePath, bytes[:number])
		if err != nil {
			ctx.String(http.StatusOK, "文件上传失败："+err.Error())
			return
		}
		//根据uid存储数据库
		fileUrl := fmt.Sprintf("http://127.0.0.1:8080/user/getImage/%s", fileName)
		err = model.ImageUpload(fileUrl, uid.(string))
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"msg": "上传成功", "Url": fileUrl})

	} else {
		//存入文件库
		corsFilePath := viper.GetString("alibaba.cors.chatImageDir") + fileName
		err := cors.UploadFile(corsFilePath, bytes[:number])
		if err != nil {
			ctx.String(http.StatusOK, "文件上传失败："+err.Error())
			return
		}
		fileUrl := fmt.Sprintf("http://127.0.0.1:8080/user/getImage/%s", fileName)
		ctx.JSON(http.StatusOK, gin.H{"msg": "上传成功", "Url": fileUrl})
	}

}

// @Summary 获取图片
// @Tags 用户模块
// @Accept json
// @Produce json
// @Param path query  string true "文件的名称"
// @Success 200 {string} json
// @router /getImage/{:path} [get]
func GetUserImage(ctx *gin.Context) {
	//uid, _ := ctx.Get("uid")
	fileName := ctx.Param("path")
	split := strings.Split(fileName, ".")
	subffix := split[len(split)-1]
	if strings.EqualFold(strings.ToLower(subffix), strings.ToLower("jpg")) ||
		strings.EqualFold(strings.ToLower(subffix), strings.ToLower("png")) {
		corsFilePath := viper.GetString("alibaba.cors.chatImageDir") + fileName
		bytes, err := API.GetCors().DownLoadFile(corsFilePath)
		if err != nil {
			ctx.String(http.StatusOK, "Err:"+err.Error())
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": bytes})
	} else {
		corsFilePath := viper.GetString("alibaba.cors.chatFileDir") + fileName
		bytes, err := API.GetCors().DownLoadFile(corsFilePath)
		if err != nil {
			ctx.String(http.StatusOK, "Err:"+err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": bytes})
	}
}

// @Summary 获取用户信息
// @Tags 用户模块
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer Token"
// @Success 200 {object} model.user
// @Router /user/getUser [post]
func GetUser(ctx *gin.Context) {
	uid, _ := ctx.Get("uid")
	user, err := model.GetUserById(uid.(string))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
