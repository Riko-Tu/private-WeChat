package API

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
	"turan.com/WeChat-Private/utils"
)

type ipInfo struct {
	Status      string
	Continent   string
	Country     string
	CountryCode string
	RegionName  string
	City        string
	Zip         string
	Lat         float32
	Lon         float32
	Timezone    string
	Currency    string
	Isp         string
	Org         string
	As          string
	Reverse     string
	Mobile      bool
	Proxy       bool
	Ip          string
}

type uvInfo struct {
	result struct {
		Uv        float32   `json:"uv"`
		UvTime    time.Time `json:"uv_time"`
		UvMax     float32   `json:"uv_max"`
		Ozone     float32   `json:"ozone"`
		OzoneTime time.Time `json:"ozone_time"`
	}
}

//获取紫外线的数据
func GetUv(lat, lng float32) ([]string, error) {
	client := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5 * time.Second,
	}
	url := fmt.Sprintf("https://api.openuv.io/api/v1/uv?lat=%0.4f&lng=%0.4f", lat, lng)
	request, err := http.NewRequest("", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("x-access-token", "8459cb92691474b6a92eeddb3cbe7be2")
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("状态码：%d", res.StatusCode)
	}
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	uvList := utils.JsonParseString(all)
	return uvList, nil

}

//获取ip地址信息
func GetIpInfo(address string) (*ipInfo, error) {
	var ip = &ipInfo{}
	fmt.Println(address)
	url := fmt.Sprintf("https://api.techniknews.net/ipgeo/%s", address)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	//使用io的实现类，多去res的body
	all, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	//将json反射到结构体
	err = json.Unmarshal(all, ip)
	if err != nil {
		return nil, err
	}
	return ip, nil
}

//每天给你一条advice
func Advice() {
	usrl2 := fmt.Sprintf("https://api.adviceslip.com/advice/%d", 2)
	get, err := http.Get(usrl2)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer get.Body.Close()
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(all))
}

// @Summary 获取图片信息
// @Tags API
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer Token"
// @Success 200 {string} string "url"
// @Failure 500 {string} json
// @Router /api/sister [get]
func GetSister(ctx *gin.Context) {
	cleint := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       5 * time.Second,
	}
	usrl2 := fmt.Sprintf("https://api.nmb.show/xiaojiejie1.php")
	get, err := http.NewRequest("get", usrl2, nil)
	get.Header = map[string][]string{"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.53"}}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	//httpcli发送请求
	do, err := cleint.Do(get)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	if do.StatusCode != 200 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	fileName := fmt.Sprintf("%v", time.Now().Unix()) + ".png"

	filePath := viper.GetString("alibaba.cors.chatImageDir") + fileName
	err = GetCors().UploadFile(filePath, all)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	fileUrl := fmt.Sprintf("http://127.0.0.1:8080/user/getImage/%s", fileName)
	ctx.String(http.StatusOK, fileUrl)
}
