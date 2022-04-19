package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
)

//编辑资料

func EditData(id int, name string, introduce string) (err error) {
	var sqlName = `UPDATE user_data set user_name = ?,user_introduce = ? where user_id = ?`
	DB := db.Get()
	DB.Exec(sqlName, name, introduce, id)

	return
}

// 获取所有未完成todo
func UserSeeTodo() (user model.User, err error) {
	var sqlName = `SELECT * FROM users`

	// 查询数据
	DB := db.Get()
	DB.Select(sqlName).Find(&user)

	return
}

func DataSeeUpdate(id int) (user model.User, err error) {
	var sqlName = `SELECT user_id,user_name,user_sex,user_birthday,user_introduce FROM user_data WHERE user_id = ?`
	// 查询数据
	DB := db.Get()
	DB.Select(sqlName, id).Find(&user)

	return
}
