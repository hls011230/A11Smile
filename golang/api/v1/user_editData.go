package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
)

//编辑资料

func EditData(uid int, uname string, resume string) (err error) {
	var sqlName = `UPDATE users set uname = ?,resume = ? where id = ?`
	DB := db.Get()
	DB.Exec(sqlName, uname, resume, uid)

	return
}

func UserSeeTodo(id int) (model.User) {
	DB := db.Get()
	var user model.User
	DB.Raw("select * from users where id =?",id).Find(&user)


	return user

}

func DataSeeUpdate(id int) (user model.User, err error) {
	var sqlName = `SELECT * FROM users WHERE id = ?`
	// 查询数据
	DB := db.Get()
	DB.Select(sqlName, id).Find(&user)

	return
}
