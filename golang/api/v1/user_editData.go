package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
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

func UserSeeTodo(id int) (model.User) {
	DB := db.Get()
	var user model.User
	DB.Raw("select * from users where id = ?",id).Find(&user)

	return user

}

func DataSeeUpdate(id int) (user model.User, err error) {
	var sqlName = `SELECT * FROM users WHERE id = ?`
	// 查询数据
	DB := db.Get()
	DB.Select(sqlName, id).Find(&user)

	return
}

//查询实名认证的数据
func UserAuthenticationSee(id *model.UserAuthentication) (authentication model.UserAuthentication) {
	var DB = db.Get()
	DB.Table("users").Select("users.resume,users.block_address,UserAuthentication.name,UserAuthentication.birthday,UserAuthentication.gender",id).Joins("left join users.id = gainer_authentication.id").Scan(&authentication)

	return


}
