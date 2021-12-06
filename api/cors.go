package API

import (
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"os"
)

var corsClient *oss.Client
var corsConfigVar *Cors

type Credential struct {
	AppKey    string
	AppSecret string
}

type Cors struct {
	Credential *Credential
	Bucket     string
	Endpoint   string
}

func CorsInfo() {

	appKey, err := base64.StdEncoding.DecodeString(viper.GetString("alibaba.accessKey"))
	if err != nil {
		panic(err.Error())
	}

	credential := &Credential{
		AppKey:    string(appKey),
		AppSecret: viper.GetString("alibaba.accessKeySecret"),
	}
	cors := &Cors{
		Credential: credential,
		Bucket:     viper.GetString("alibaba.cors.bucket"),
		Endpoint:   viper.GetString("alibaba.cors.Endpoint"),
	}

	corsConfigVar = cors
	clientInfo(corsConfigVar)

}

func GetCors() *Cors {
	return corsConfigVar
}
func clientInfo(c *Cors) {

	client, err := oss.New(c.Endpoint, c.Credential.AppKey, c.Credential.AppSecret)
	if err != nil {
		panic(err.Error())
	}
	corsClient = client
}

// 创建存储桶
func (c Cors) CreateBucket() {
	//创建标注存储空间
	err := corsClient.CreateBucket(c.Bucket, oss.ACL(oss.ACLPublicRead))
	if err != nil {
		panic(err.Error())
	}
	//同城区域冗余存储
	//err = c.getClient().CreateBucket(bucketName, oss.RedundancyType(oss.RedundancyZRS))
	//if err != nil {
	//	panic(err.Error())
	//}
}

// 查询所有存储桶
func (c Cors) GetDucketList() {
	marker := ""

	lsRes, err := corsClient.ListBuckets(oss.Marker(marker))
	if err != nil {
		panic(err.Error())
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("bucket:", bucket.Name)
	}

}

// 上传文件

// 下载文件

// 获取路径下文件列表

// 获取临时秘钥，指定使用的cosAPI和cos文件路径

// 设置跨域配置
func (c Cors) SetCorsConfig() {
	rule1 := oss.CORSRule{
		AllowedOrigin: []string{"*"},
		AllowedMethod: []string{"PUT", "GET", "POST", "DELETE"},
		AllowedHeader: []string{"Authorization"},
		ExposeHeader:  []string{},
		MaxAgeSeconds: 200,
	}

	//rule2 := oss.CORSRule{
	//	AllowedOrigin: []string{"http://example.com", "http://example.net"},
	//	AllowedMethod: []string{"POST"},
	//	AllowedHeader: []string{"Authorization"},
	//	ExposeHeader:  []string{"x-oss-test", "x-oss-test1"},
	//	MaxAgeSeconds: 100,
	//}
	err := corsClient.SetBucketCORS(c.Bucket, []oss.CORSRule{rule1})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

//获取配置信息
func (c Cors) GetCorsConfig() {
	corsRes, err := corsClient.GetBucketCORS(c.Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("Bucket CORS:", corsRes.CORSRules)
}
