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
