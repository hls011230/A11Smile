package v1

import (
	"A11Smile/db/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)


func PostMedicalInformation(f io.Reader, token model.RespWXToken, uid int) error {

	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)
	ContentType := writer.FormDataContentType()

	// 提供表单中的字段名<img>和文件名<new.jpg>,返回值是可写的接口io.Writer
	formFile, err1 := writer.CreateFormFile("img", "new.jpg")
	if err1 != nil {
		fmt.Println("创建form文件失败:", err1)
		return err1
	}


	_, err := io.Copy(formFile, f)
	if err != nil {
		fmt.Println("写入form文件失败", err)
		return err
	}

	writer.Close()

	// POST请求调用微信身份证OCR接口
	url := "https://api.weixin.qq.com/cv/ocr/comm?access_token=%s"
	res, err2 := http.Post(fmt.Sprintf(url, token.Access_token), ContentType, buf)
	if err2 != nil {
		fmt.Println("OCR接口请求失败:", err2)
		return err2
	}

	defer res.Body.Close()

	// Json数据绑定返回数据包
	var medicalInformation model.RespWXMedicalInformation
	err = json.NewDecoder(res.Body).Decode(&medicalInformation)
	fmt.Println(medicalInformation.Items[0].Text)
	for k, v := range medicalInformation.Items {
		fmt.Printf("第%v个值为:%v\n",k,v)
	}
	return nil
}
