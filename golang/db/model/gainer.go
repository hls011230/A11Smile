package model

type GainerAuthentication struct {
	Id             int    `json:"id" gorm:"column:id"`
	RegNum         string `json:"reg_num" gorm:"column:reg_num"`
	Address         string `json:"address" gorm:"column:address"`
	EnterpriseName string `json:"enterprise_name" gorm:"column:enterprise_name"`
}

type Gainer struct {
	Id              int    `json:"id" gorm:"column:id"`
	Gid             int    `json:"gid" gorm:"column:gid"`
	Email           string `json:"email" gorm:"column:email"`
	Passwd          string `json:"passwd" gorm:"column:passwd"`
	Resume          string `json:"resume" gorm:"column:resume"`
	BlockAddress    string `json:"block_address" gorm:"column:block_address"`
	PrivateKey      string `json:"private_key" gorm:"column:private_key"`
	ImgUrl string `json:"img_url" gorm:"column:img_url"`
}
