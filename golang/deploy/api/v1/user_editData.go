package v1

import (
	"A11Smile/deploy/db/model"
	"database/sql"
)

//编辑资料
var DB *sql.DB
func init() {
	var err error
	var source = "root:123456@tcp(localhost:3306)/a11smile?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err =  sql.Open("mysql",source)
	if err != nil {
		return
	}
}

func EditData(id int,name string,introduce string) (err error) {
	var sqlName = `UPDATE user_data set user_name = ?,user_introduce = ? where user_id = ?`
	_,err = DB.Exec(sqlName,name,introduce,id)
	if err != nil{
		return err
	}


	return
}

// 获取所有未完成todo
func UserSeeTodo() (todos []model.User, err error) {
	var sqlName = `SELECT user_id,user_name,user_sex,user_birthday,user_introduce FROM user_data`

	// 查询数据
	rows, err := DB.Query(sqlName)
	if err != nil {
		return nil, err
	}

	// 按行读取数据到结构体成员中
	for rows.Next() {
		todo := model.User{}
		err = rows.Scan(&todo.Id, &todo.Name, &todo.Sex, &todo.Birthday,&todo.Introduce)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return
}

func DataSeeUpdate(id int) (todos []model.User, err error) {
	var sqlName = `SELECT user_id,user_name,user_sex,user_birthday,user_introduce FROM user_data WHERE user_id = ?`
	// 查询数据
	rows, err := DB.Query(sqlName,id)
	if err != nil {
		return nil, err
	}

	// 按行读取数据到结构体成员中
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Sex, &user.Birthday,&user.Introduce)
		if err != nil {
			return nil, err
		}
		todos = append(todos, user)
	}

	return
}
