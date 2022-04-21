package model

// 定义数据库访问模型
type User struct {
	Id           int    `json:"id" gorm:"column:id"`
	Email        string `json:"email" gorm:"column:email"`
	Uname        string `json:"uname" gorm:"column:uname"`
	Passwd       string `json:"passwd" gorm:"column:passwd"`
	Resume       string `json:"resume" gorm:"column:resume"`
	BlockAddress string `json:"block_address" gorm:"column:block_address"`
	PrivateKey   string `json:"private_key" gorm:"column:private_key"`
}

// 定义用户真实身份模型
type UserAuthentication struct {
	Id          int    `json:"id" gorm:"column:id"`
	Uid         int    `json:"uid" gorm:"column:uid"`
	Name        string `json:"name" gorm:"column:name"`
	Birthday    string `json:"birthday" gorm:"column:birthday"`
	Nationality string `json:"nationality" gorm:"column:nationality"`
	Gender      string `json:"gender" gorm:"column:gender"`
}

// 登录用户模型
type LoginUser struct {
	Email    string `json:"email"`
	Passwd string `json:"passwd"`
}

// 使用合约User对象

type UserWallet struct {
	BlockAddress string `json:"block_address" gorm:"column:block_address"`
	PrivateKey   string `json:"private_key" gorm:"column:private_key"`
}
