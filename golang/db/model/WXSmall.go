package model

// 凭证对象
type RespWXToken struct {
	Access_token string `json:"access_token"` // 获取到的凭证
	Expires_in   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode      int    `json:"errcode"`      // 错误码
	ErrMsg       string `json:"errMsg"`       // 错误信息
}

// 身份证识别对象
type RespWXIDOCRF struct {
	ErrCode     int    `json:"errCode"`     // 错误码
	ErrMsg      string `json:"errMsg"`      // 错误信息
	Id          string `json:"id"`          // 身份证号
	Name        string `json:"name"`        // 名字
	Gender      string `json:"gender"`      // 性别
	Nationality string `json:"nationality"` // 民族
}

// 小程序AppID和密钥
type wxApp struct {
	AppId  string
	Secret string
}

var WXAPP = wxApp{
	AppId:  "wx1f6551a9b737be68",
	Secret: "e8c9ec131ccbcc684b6cd0327404c869",
}

// 营业执照识别对象
type RespWXBizlicense struct {
	ErrCode            int    `json:"errCode"`              // 错误码
	ErrMsg             string `json:"errMsg"`               // 错误信息
	RegNum             string `json:"reg_num"`              // 企业注册号
	EnterpriseName     string `json:"enterprise_name"`      // 企业名称
	TypeOfOrganization string `json:"type_of_organization"` // 企业类型
	RegisteredDate     string `json:"registered_date"`      // 企业注册时间
}

// 医疗信息识别对象

type RespWXMedicalInformation struct {
	ErrCode int    `json:"errCode"` // 错误码
	ErrMsg  string `json:"errMsg"`  // 错误信息
	Items   []Item `json:"items"`   //识别结果
}

type Item struct {
	Text string      `json:"text"`
	Pos  interface{} `json:"pos"`
}

// 文件上传链接
type RespWXUploadLink struct {
	ErrCode       int    `json:"errCode"` // 错误码
	ErrMsg        string `json:"errMsg"`  // 错误信息
	Url           string `json:"url"`
	Token         string `json:"token"`
	Authorization string `json:"authorization"`
	FileId        string `json:"file_id"`
	CosFileId     string `json:"cos_file_id"`
}

// 文件对象
type UserCloudLink struct {
	FileId string `json:"fileid"`
	MaxAge int `json:"max_age"`
}

type RespWXLoadLink struct {
	ErrCode       int    `json:"errCode"` // 错误码
	ErrMsg        string `json:"errMsg"`  // 错误信息
	FileList   []UserLoadUrl `json:"file_list"`
}

type UserLoadUrl struct {
	DownloadUrl string `json:"download_url"`
	Status int `json:"status"`
}

