package v1

import (
	"A11Smile/db"
	"errors"
)

//编辑用户名
func EditUserName(uid int, uname string) (err error) {
	var sqlName = `UPDATE users set uname = ? where id = ?`
	DB := db.Get()
	DB.Exec(sqlName, uname, uid)

	return
}
//修改用户简介
func EditUserResume(uid int, resume string) (err error) {
	var sqlName = `UPDATE users set resume = ? where id = ?`
	DB := db.Get()
	DB.Exec(sqlName, resume, uid)

	return
}

//查询实名认证的数据
func UserAuthenticationSee(id int) (interface{},error){
	user := struct {
		Block_address string `json:"block_address"`
		Birthday      string `json:"birthday"`
		Resume        string `json:"resume"`
		Uname         string `json:"uname"`
		Gender        string `json:"gender"`
	}{}
	var DB = db.Get()
	DB.Table("users").Select("users.resume,users.block_address,users.uname,user_authentication.birthday,user_authentication.gender").Joins("left join user_authentication on user_authentication.uid = users.id where users.id = ?",id).Scan(&user)
	if user.Gender == "" {
		return nil,errors.New("还未进行实名认证")
	}
	return user,nil


}
