package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
)

//编辑
func GainerEditData(gid int, resume string) (err error) {
	var sqlName = `UPDATE gainer_authentication set resume = ? where id = ?`
	DB := db.Get()
	DB.Exec(sqlName, resume, gid)

	return
}

// 获取所有未完成todo
func GainerSeeTodo(id int) (model.GainerAuthentication) {
	DB := db.Get()
	var gainer model.GainerAuthentication
	DB.Raw("select * from gainer_authentication where id =?",id).Find(&gainer)


	return gainer
}

func GainerDataSeeUpdate(gid int) (gainer model.GainerAuthentication, err error) {
	var sqlName = `SELECT id,gid,reg_num,serial,enterprise_name,introduce FROM gainer_authentication WHERE gid = ?`
	// 查询数据
	DB := db.Get()
	DB.Select(sqlName,gid).Find(&gainer)

	return
}
