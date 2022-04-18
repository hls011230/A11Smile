package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"bytes"
	"encoding/json"
	"fmt"

	"io"
	"mime/multipart"
	"net/http"
)

func PostIDCard(f io.Reader, token model.RespWXToken) error {

	buf := new(bytes.Buffer)                    // 实例化一个结构体
	writer := multipart.NewWriter(buf)          // 返回一个writer指针
	ContentType := writer.FormDataContentType() // 表单提交的格式

	// 提供表单中的字段名<img>和文件名<new.jpg>,返回值是可写的接口io.Writer
	formFile, err := writer.CreateFormFile("img", "new.jpg")
	if err != nil {
		fmt.Println("创建form文件失败:", err)
		return err
	}

	_, err = io.Copy(formFile, f)
	if err != nil {
		fmt.Println("写入form文件失败", err)
		return err
	}

	writer.Close()

	// POST请求调用微信身份证OCR接口
	url := "https://api.weixin.qq.com/cv/ocr/idcard?type=photo&access_token=%s"
	res, err2 := http.Post(fmt.Sprintf(url, token.Access_token), ContentType, buf)
	if err2 != nil {
		fmt.Println("OCR接口请求失败:", err2)
		return err2
	}

	defer res.Body.Close()

	//Json数据绑定返回数据包
	var Front_IDCard model.RespWXIDOCRF
	var UserAuthentication model.UserAuthentication
	err = json.NewDecoder(res.Body).Decode(&Front_IDCard)
	if err != nil {
		fmt.Println("数据绑定失败:", err)
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&UserAuthentication)
	if err != nil {
		fmt.Println("数据绑定失败:", err)
		return err
	}

	cli := db.Get()
	cli.Table("user_authentication").Save(UserAuthentication)

	return nil
}
