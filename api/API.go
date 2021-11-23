package API

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

//获取ip地址信息
func GetIpInfo(ip string) {
	url := fmt.Sprintf("https://api.techniknews.net/ipgeo/%s", ip)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Body.Close()
	all, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(all))
	//将string转成json 取值

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
		ctx.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	//httpcli发送请求
	do, err := cleint.Do(get)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	if do.StatusCode != 200 {
		ctx.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	all, err := ioutil.ReadAll(do.Body)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, all)
}
