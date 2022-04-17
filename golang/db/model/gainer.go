package model

type GainerAuthentication struct {
	Id             int    `json:"id" gorm:"column:id"`
	RegNum         string `json:"reg_num" gorm:"column:reg_num"`
	Serial         string `json:"serial" gorm:"column:serial"`
	EnterpriseName string `json:"enterprise_name" gorm:"column:enterprise_name"`
}
