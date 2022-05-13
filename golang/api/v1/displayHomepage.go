package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net/http"
)

func DisplayHomepage()([]interface{},error) {
	r1,err := eth.Ins.SeeGainerMedicalInformationsName(nil)
	if err != nil {
		return nil, err
	}
	DB := db.Get()
	var res []interface{}

	for _, v := range r1 {


		var gainer model.Gainer
		DB.Table("gainers").First(&gainer,"block_address = ?",v.Addr.String())

		// 查询指定病历文件的云地址
		var fileCloudPath string
		DB.Raw("select img_url from gainers where id = ?",gainer.Id).Find(&fileCloudPath)

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
			return res,err
		}

		req.Header.Set("Content-Type", "application/json")
		resp, _ := http.DefaultClient.Do(req)

		// Json数据绑定
		var respWXLoadLink model.RespWXLoadLink
		err = json.NewDecoder(resp.Body).Decode(&respWXLoadLink)
		if err != nil {
			return res,err
		}

		r := struct {
			MedicalName  string
			Min          *big.Int
			Max          *big.Int
			Addr         common.Address
			HospitalName string
			Account      *big.Int
			Exit         bool
			IconUrl  string
			Department string
		}{
			MedicalName : v.MedicalName,
			Min          :v.Min,
			Max         :v.Max,
			Addr        :v.Addr,
			HospitalName :v.HospitalName,
			Account     :v.Account,
			Exit         :v.Exit,
			IconUrl: respWXLoadLink.FileList[0].DownloadUrl,
			Department: v.Department,
		}
		res = append(res, r)
	}


	return res,nil
}

func ShowDetailsPage(details model.PostDetails) (interface{},error) {
	res,err := eth.Ins.SeeGainerMedicalInformations(nil,common.HexToAddress(details.Address),details.MedicalName)
	if err != nil {
		return res, err
	}

	DB := db.Get()
	var gainer model.Gainer
	DB.Table("gainers").First(&gainer,"block_address = ?",details.Address)

	// 查询指定病历文件的云地址
	var fileCloudPath string
	DB.Raw("select img_url from gainers where id = ?",gainer.Id).Find(&fileCloudPath)

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
		return res,err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)

	// Json数据绑定
	var respWXLoadLink model.RespWXLoadLink
	err = json.NewDecoder(resp.Body).Decode(&respWXLoadLink)
	if err != nil {
		return res,err
	}

	DetailsPage := struct {
		Min                       *big.Int
		Max                       *big.Int
		Medicalrecordrequirements string
		Requirementdescription    string
		Resume string
		HospitalName string
		IconUrl string
	}{
		Min: res.Min,
		Max: res.Max,
		Medicalrecordrequirements: res.Medicalrecordrequirements,
		Requirementdescription: res.Requirementdescription,
		Resume: gainer.Resume,
		HospitalName: details.HospitalName,
		IconUrl:respWXLoadLink.FileList[0].DownloadUrl,
	}

	return DetailsPage,nil
}

func ShowSortPage(sort string) ([]interface{},error) {
	r1,err := eth.Ins.SeeGainerMedicalInformationsName(nil)
	if err != nil {
		return nil, err
	}
	DB := db.Get()
	var res []interface{}

	for _, v := range r1 {
		if v.Department == sort {
			var gainer model.Gainer
			DB.Table("gainers").First(&gainer,"block_address = ?",v.Addr.String())

			// 查询指定病历文件的云地址
			var fileCloudPath string
			DB.Raw("select img_url from gainers where id = ?",gainer.Id).Find(&fileCloudPath)

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
				return res,err
			}

			req.Header.Set("Content-Type", "application/json")
			resp, _ := http.DefaultClient.Do(req)

			// Json数据绑定
			var respWXLoadLink model.RespWXLoadLink
			err = json.NewDecoder(resp.Body).Decode(&respWXLoadLink)
			if err != nil {
				return res,err
			}

			r := struct {
				MedicalName  string
				Min          *big.Int
				Max          *big.Int
				Addr         common.Address
				HospitalName string
				Account      *big.Int
				Exit         bool
				IconUrl  string
				Department string
			}{
				MedicalName : v.MedicalName,
				Min          :v.Min,
				Max         :v.Max,
				Addr        :v.Addr,
				HospitalName :v.HospitalName,
				Account     :v.Account,
				Exit         :v.Exit,
				IconUrl: respWXLoadLink.FileList[0].DownloadUrl,
				Department: v.Department,
			}
			res = append(res, r)
		}
	}


	return res,nil
}