package utils

import (
	"encoding/json"
	"fmt"
	"github.com/PaesslerAG/jsonpath"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gopkg.in/gomail.v2"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

//邮箱正则
var validEmail = regexp.MustCompilePOSIX("^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$")

//邮箱校验
func VerifyEmail(email string) bool {
	return validEmail.MatchString(email)
}

//获取验证码
func GetCode() string {
	//获取活种
	nano := time.Now().UnixNano()
	//每个种子对应一个随机值
	rnd := rand.New(rand.NewSource(nano))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

var (
	from     = viper.GetString("smtp.from")
	host     = viper.GetString("smtp.qq.host")
	port, _  = strconv.Atoi(viper.GetString("smtp.qq.port"))
	username = viper.GetString("smtp.qq.username")
	password = viper.GetString("smtp.qq.password")
)

//发送邮件
func SendEmail(addressEmail string, code string) error {
	m := gomail.NewMessage()                   //获取邮件对象
	m.SetHeader("From", "wechat"+"<"+from+">") //发件人邮箱
	m.SetHeader("To", addressEmail)            //收件人邮箱
	m.SetHeader("Subject", "chat【验证码】")        //标题
	m.SetBody("text/html", fmt.Sprintf("你的验证码是%s", code))

	//创建smtp拨号器
	d := gomail.Dialer{Host: host, Port: port, Username: username, Password: password}
	//使用拨号器发送message
	return d.DialAndSend(m)
}

//uuid生成
func GetUuid(email string) uuid.UUID {
	//随机生成一个UUID
	v1 := uuid.NewV1()
	//将随机的uuid与邮箱结合
	v3 := uuid.NewV3(v1, email)
	return v3
}

//jsonPath

func JsonParseString(jsonString []byte) []string {
	//[]byte(`{"result":{
	//								"uv":6.3019,"uv_time":"2021-11-26T02:21:04.362Z",
	//								"uv_max":15.5084,
	//								"uv_max_time":"2021-11-25T23:02:51.036Z",
	//								"ozone":246.7,"ozone_time":"2021-11-26T00:04:12.661Z",
	//								"safe_exposure_time":{"st1":26,"st2":32,"st3":42,"st4":53,"st5":85,"st6":159},
	//								"sun_info":{
	//								"sun_times":{
	//								"solarNoon":"2021-11-25T23:02:51.036Z",
	//								"nadir":"2021-11-25T11:02:51.036Z",
	//								"sunrise":"2021-11-25T16:07:01.627Z",
	//								"sunset":"2021-11-26T05:58:40.446Z",
	//								"sunriseEnd":"2021-11-25T16:09:44.841Z",
	//								"sunsetStart":"2021-11-26T05:55:57.231Z",
	//								"dawn":"2021-11-25T15:40:13.359Z",
	//								"dusk":"2021-11-26T06:25:28.713Z",
	//								"nauticalDawn":"2021-11-25T15:07:51.188Z",
	//								"nauticalDusk":"2021-11-26T06:57:50.884Z",
	//								"nightEnd":"2021-11-25T14:33:32.598Z",
	//								"night":"2021-11-26T07:32:09.475Z",
	//								"goldenHourEnd":"2021-11-25T16:41:23.177Z",
	//								"goldenHour":"2021-11-26T05:24:18.895Z"},
	//								"sun_position":{
	//								"azimuth":1.5665669112083027,
	//								"altitude":0.77600758347787}}}}`)
	var v = interface{}(nil)
	zap.L().Debug(string(jsonString))
	err := json.Unmarshal(jsonString, &v)
	if err != nil {
		fmt.Println(err.Error())
	}
	result := []string{`$.result.uv`, `$.result.uv_time`, `$.result.uv_max`, `$.result.uv_max_time`}
	uvInfo := make([]string, 0, 10)
	for i := 0; i < len(result); i++ {
		value, err := jsonpath.Get(result[i], v)
		if err != nil {
			fmt.Println(123, err.Error())
		}
		valueString := fmt.Sprintf("%v", value)
		uvInfo = append(uvInfo, valueString)
	}
	zap.L().Debug(fmt.Sprintf("%v", uvInfo))
	return uvInfo
}
