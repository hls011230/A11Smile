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
)

func PostBizlicense(f io.Reader, token model.RespWXToken) error {

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
		return err
	}

	writer.Close()

	// POST请求调用微信营业执照OCR接口
	url := "https://api.weixin.qq.com/cv/ocr/bizlicense?access_token=%s"
	res, err2 := http.Post(fmt.Sprintf(url, token.Access_token), ContentType, buf)
	if err2 != nil {
		fmt.Println("OCR接口请求失败:", err2)
		return err2
	}

	defer res.Body.Close()

	// Json数据绑定返回数据包
	var gainer_authentication model.GainerAuthentication

	err = json.NewDecoder(res.Body).Decode(&gainer_authentication)
	if err != nil {
		fmt.Println("数据绑定失败:", err)
		return err
	}

	// 判断营业执照是否合格
	if gainer_authentication.RegNum == "" {
		return errors.New("认证失败,请确认您上传的营业执照正确")
	}

	// 保存至数据库
	cli := db.Get()
	cli.Table("gainer_authentication").Save(gainer_authentication)

	return nil
}
