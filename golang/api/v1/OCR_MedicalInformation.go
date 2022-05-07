package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)


func PostMedicalInformation(f io.Reader, token model.RespWXToken, uid int) error {

	DB := db.Get()
	var name string
	DB.Table("users").Select("user_authentication.name").Joins("left join user_authentication on user_authentication.uid = users.id where users.id = ?",uid).Scan(&name)

	if name == "" {
		return errors.New("请先进行实名认证")
	}

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
	count := 0
	for _, v := range medicalInformation.Items {
		if strings.Contains(v.Text, name) {
			count += 1
		}
	}
	if count != 1 {
		return errors.New("请上传真实信息")
	}
	return nil
}
