package v1

import (
	"A11Smile/deploy/db/model"
	"database/sql"
)

//编辑资料
var gainerDB *sql.DB
func init() {
	var err error
	var source = "root:123456@tcp(localhost:3306)/a11smile?charset=utf8mb4&parseTime=True&loc=Local"
	gainerDB, err =  sql.Open("mysql",source)
	if err != nil {
		return
	}
}

func GainerEditData(gid int,enterprise_name string,introduce string) (err error) {
	var sqlName = `UPDATE gainer_authentication set enterprise_name = ?,introduce = ? where gid = ?`
	_,err = gainerDB.Exec(sqlName,enterprise_name,introduce,gid)
	if err != nil{
		return err
	}


	return
}

// 获取所有未完成todo
func GainerSeeTodo() (todos []model.Gainer, err error) {
	var sqlName = `SELECT id,gid,reg_num,serial,enterprise_name,introduce FROM gainer_authentication`

	// 查询数据
	rows, err := gainerDB.Query(sqlName)
	if err != nil {
		return nil, err
	}

	// 按行读取数据到结构体成员中
	for rows.Next() {
		gainer := model.Gainer{}
		err = rows.Scan(&gainer.Id, &gainer.Gid, &gainer.Reg_num, &gainer.Enterprise_name,&gainer.Introduce)
		if err != nil {
			return nil, err
		}
		todos = append(todos, gainer)
	}

	return
}

func GainerDataSeeUpdate(gid int) (todos []model.Gainer, err error) {
	var sqlName = `SELECT id,gid,reg_num,serial,enterprise_name,introduce FROM gainer_authentication WHERE gid = ?`
	// 查询数据
	rows, err := gainerDB.Query(sqlName,gid)
	if err != nil {
		return nil, err
	}

	// 按行读取数据到结构体成员中
	for rows.Next() {
		gainer := model.Gainer{}
		err = rows.Scan(&gainer.Id, &gainer.Gid, &gainer.Reg_num,&gainer.Serial, &gainer.Enterprise_name,&gainer.Introduce)
		if err != nil {
			return nil, err
		}
		todos = append(todos,gainer )
	}

	return
}
