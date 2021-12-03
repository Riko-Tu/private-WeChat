package API

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Credential struct {
	AppKey    string
	AppSecret string
}

type Cors struct {
	Credential *Credential
	bucket     string
	Endpoint   string
}

// 获取cos client
func (c Cors) getClient() *oss.Client {
	//url, _ := url2.Parse(fmt.Sprintf("https:%v.%v", c.bucket, c.Endpoint))
	client, err := oss.New(c.Endpoint, c.Credential.AppKey, c.Credential.AppSecret)
	if err != nil {
		panic(err.Error())
	}
	return client
}

// 创建存储桶
func (c Cors) CreateBucket(bucketName string) {
	//创建标注存储空间
	err := c.getClient().CreateBucket(bucketName, oss.ACL(oss.ACLPublicRead))
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
	for {
		lsRes, err := c.getClient().ListBuckets(oss.Marker(marker))
		if err != nil {
			panic(err.Error())
		}
		for _, bucket := range lsRes.Buckets {
			fmt.Println("bucket:", bucket.Name)
		}
	}

}

// 上传文件

// 下载文件

// 获取路径下文件列表

// 获取临时秘钥，指定使用的cosAPI和cos文件路径

// 设置跨域配置
