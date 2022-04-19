package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"

	"io"
	"mime/multipart"
	"net/http"
)

func PostIDCard(f io.Reader, token model.RespWXToken, uid int) error {

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

	year, month := GetAge(Front_IDCard.Id)

	UserAuthentication.Uid = uid
	UserAuthentication.Name = Front_IDCard.Name
	UserAuthentication.Nationality = Front_IDCard.Nationality
	UserAuthentication.Gender = Front_IDCard.Gender
	UserAuthentication.Birthday = fmt.Sprintf("%v年%v月", year, month)

	cli := db.Get()

	cli.Exec("insert into user_authentication (`uid`,`name`,`birthday`,`nationality`,`gender`) VALUES (?,?,?,?,?)", UserAuthentication.Uid, UserAuthentication.Name, UserAuthentication.Birthday, UserAuthentication.Nationality, UserAuthentication.Gender)

	return nil
}

func GetAge(identification_number string) (int, int) {
	reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)(\d{2})((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	//reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)`)
	params := reg.FindStringSubmatch(identification_number)
	birYear, _ := strconv.Atoi(params[1] + params[2])
	birMonth, _ := strconv.Atoi(params[3])

	return birYear, birMonth
}
