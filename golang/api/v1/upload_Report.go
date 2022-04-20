package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

func UploadMedicalHistory(srcFile io.Reader, token model.RespWXToken, uid int) error {

	// 获取用户的Address值
	var addr string
	DB := db.Get()
	DB.Raw("select block_address from users where id = ?",uid).Find(&addr)

	// 获取文件上传地址
	fileName := time.Now().Unix()
	path := fmt.Sprintf("a11smile/users/%v/%v.jpg",addr,fileName)
	myReq := struct {
		Env  string `json:"env"`
		Path string `json:"path"`
	}{
		Env:  "prod-9gy59jvo10e0946b",
		Path: path,
	}

	reqByte, err := json.Marshal(myReq)

	u := "https://api.weixin.qq.com/tcb/uploadfile?access_token=%s"

	req, err := http.NewRequest("POST", fmt.Sprintf(u, token.Access_token), bytes.NewReader(reqByte))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	cli := http.Client{
		Timeout: 45 * time.Second,
	}

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	// Json数据绑定
	var respUploadLink model.RespWXUploadLink
	err = json.NewDecoder(resp.Body).Decode(&respUploadLink)
	if err != nil {
		return err
	}

	fmt.Println(respUploadLink)
	// 上传文件

	myUploadReq := struct {
		Key               string    `json:"key" form:"key"`
		Signature         string    `json:"Signature" form:"Signature"`
		XCosSecurityToken string    `json:"x_cos_security_token" form:"x_cos_security_token"`
		XCosMetaFileid    string    `json:"x_cos_meta_fileid" form:"x_cos_meta_fileid"`
		File              io.Reader `json:"file" form:"file"`
	}{
		Key:               path,
		Signature:         respUploadLink.Authorization,
		XCosSecurityToken: respUploadLink.Token,
		XCosMetaFileid:    respUploadLink.CosFileId,
		File:              srcFile,
	}

	//读出文本文件数据

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	content_type := w.FormDataContentType()


	_ = w.WriteField("key", myUploadReq.Key)
	_ = w.WriteField("Signature", myUploadReq.Signature)
	_ = w.WriteField("x_cos_security_token", myUploadReq.XCosSecurityToken)
	_ = w.WriteField("x_cos_meta_fileid", myUploadReq.XCosMetaFileid)

	//将文件数据写入
	formFile, _ := w.CreateFormFile("file", "new.jpg")
	_, err = io.Copy(formFile, myUploadReq.File)

	_ = w.Close()

	req, err = http.NewRequest("POST", respUploadLink.Url, buf)
	if err != nil {
		fmt.Println("req", err)
		return err
	}

	req.Header.Set("Content-Type", content_type)
	resp, _ = http.DefaultClient.Do(req)



	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return err
	}
	fmt.Println(string(b))

	return nil

}


