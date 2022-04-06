package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var secretid string = "wx1f6551a9b737be68"
var secretkey string = "e8c9ec131ccbcc684b6cd0327404c869"
var cos_url string ="https://prod-9gy59jvo10e0946b-1310014865.ap-shanghai.app.tcloudbase.com/"

func user_uploadTencentFileHandle(c *gin.Context)  {
	u, _ := url.Parse("cos_url")
	b := &cos.BaseURL{BucketURL: u}
	a := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretid,
			SecretKey: secretkey,
		},
	})

	name := "体检报告/file.txt"
	// 1. 通过字符串上传对象
	f := strings.NewReader("体检报告")
	//
	_,err := a.Object.Put(context.Background(),name,f,nil)
	if err != nil{
		panic(err)
	}

	// 上传本地文件
	fileName := time.Now().Unix()
	fildDir := fmt.Sprintf("%s%d%s/",time.Now().Year(),time.Now().Month().String())
	filePath := fmt.Sprintf("%s%s%s", fildDir, fileName)
	_, err = a.Object.PutFromFile(context.Background(),name,filePath,nil)
	if err != nil {
		panic(err)
	}

	////下载文件
	//	// 1.通过响应体获取对象
	//	resp,err := a.Object.Get(context.Background(),name,nil)
	//	if err != nil{
	//		panic(err)
	//	}
	//
	//	bs,_ := ioutil.ReadAll(resp.Body)
	//	resp.Body.Close()
	//	fmt.Println("%s\n",string(bs))
	//
	//	// 2.获取对象到本地文件
	//	_,err = a.Object.GetToFile(context.Background(),name,"exampleobject",nil)
	//	if err != nil{
	//		panic(err)
	//	}
	}


