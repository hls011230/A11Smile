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
	"time"
)

// 编辑征求者资料
func GainerEditData(gid int, resume string) (err error) {
	var sqlName = `UPDATE gainers set resume = ? where id = ?`
	DB := db.Get()
	DB.Exec(sqlName, resume, gid)

	return
}

// 查询企业信息
func GainerAuthenticationSee(id int) interface{}{
	gainer := struct {
		EnterpriseName string `json:"enterprise_name"`
		Block_address  string `json:"block_address"`
		Resume         string `json:"resume"`
	}{}
	DB := db.Get()
	DB.Table("gainers").Select("gainer_authentication.enterprise_name,gainers.block_address,gainers.resume").Joins("left join gainer_authentication on gainer_authentication.id = gainers.gid where gainers.id = ?",id).Scan(&gainer)

	return gainer
}


// 更换征求者头像
func EditGainerIcon(gid int , srcFile io.Reader) error {
	DB := db.Get()
	var user model.Wallet
	DB.Table("gainers").First(&user,"id = ?" , gid)

	// 获取文件上传地址
	fileName := time.Now().Unix()
	path := fmt.Sprintf("a11smile/gainers/%v/Image/%v.jpg",user.BlockAddress,fileName)
	myReq := struct {
		Env  string `json:"env"`
		Path string `json:"path"`
	}{
		Env:  "prod-9gy59jvo10e0946b",
		Path: path,
	}

	reqByte, err := json.Marshal(myReq)

	token,_ := GetToken()
	u := "https://api.weixin.qq.com/tcb/uploadfile?access_token=%s"

	req, err := http.NewRequest("POST", fmt.Sprintf(u, token.Access_token), bytes.NewReader(reqByte))
	if err != nil {
		fmt.Println(1)
		fmt.Println(err)
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	cli := http.Client{
		Timeout: 45 * time.Second,
	}

	resp, err := cli.Do(req)
	if err != nil {
		fmt.Println(2)
		fmt.Println(err)
		return err
	}

	// Json数据绑定
	var respUploadLink model.RespWXUploadLink
	err = json.NewDecoder(resp.Body).Decode(&respUploadLink)
	if err != nil {
		return err
	}

	DB.Exec("update gainers set img_url = ? where id = ?",respUploadLink.FileId,gid)

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
	_ = w.WriteField("x-cos-security_token", myUploadReq.XCosSecurityToken)
	_ = w.WriteField("x-cos-meta-fileid", myUploadReq.XCosMetaFileid)

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


	return nil
}

func ShowGainerIcon(gid int) (string,error) {
	// 获取对象钱包
	var user model.Wallet
	DB := db.Get()
	DB.Table("gainers").First(&user,"id = ?",gid)

	// 查询指定病历文件的云地址
	var fileCloudPath string
	DB.Raw("select img_url from gainers where id = ?",gid).Find(&fileCloudPath)

	// 根据云地址获取下载地址
	token,_ := GetToken()
	cloudFile := model.UserCloudLink{FileId: fileCloudPath,MaxAge: 720}
	fileList := []model.UserCloudLink{cloudFile}

	myReq := struct {
		Env  string `json:"env"`
		FileList []model.UserCloudLink `json:"file_list"`
	}{
		Env:  "prod-9gy59jvo10e0946b",
		FileList: fileList,
	}

	reqByte, err := json.Marshal(myReq)

	url := "https://api.weixin.qq.com/tcb/batchdownloadfile?access_token=%s"
	req, err := http.NewRequest("POST", fmt.Sprintf(url, token.Access_token), bytes.NewReader(reqByte))
	if err != nil {
		return "",err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)

	// Json数据绑定
	var respWXLoadLink model.RespWXLoadLink
	err = json.NewDecoder(resp.Body).Decode(&respWXLoadLink)
	if err != nil {
		return "",err
	}


	return respWXLoadLink.FileList[0].DownloadUrl, nil
}
