package model

// 定义数据库访问模型
type User struct {
	Id            int    `json:"id" gorm:"column:id`
	Email         string `json:"email" gorm:"column:email"`
	Password      string `json:"password" gorm:"column:password"`
	Name          string `json:"name" gorm:"column:name"`
	Sex           string `json:"sex" gorm:"column:sex"`
	Birthday      string `json:"birthday" gorm:"column:birthday"`
	User_type     int    `json:"user_type" gorm:"column:user_type"`
	Block_address string `json:"block_address" gorm:"column:block_address"`
	Key_store     string `json:"key_store" gorm:"column:key_store"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RespWXSmall struct {
	Access_token string `json:"access_token"` // 获取到的凭证
	Expires_in   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode      int    `json:"errcode"`      // 错误码
	ErrMsg       string `json:"errMsg"`       // 错误信息
}

type RespWXIDOCRF struct {
	ErrCode     int    `json:"errCode"`     // 错误码
	ErrMsg      string `json:"errMsg"`      // 错误信息
	Id          string `json:"id"`          // 身份证号
	Name        string `json:"name"`        // 名字
	Gender      string `json:"gender"`      // 性别
	Nationality string `json:"nationality"` // 民族
}
