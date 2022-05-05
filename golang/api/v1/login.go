package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
)

func UserLogin(user *model.LoginUser) int {
	cli := db.Get()
	var UserId int
	cli.Raw("select id from users where email= ? and passwd = ?", user.Email, user.Passwd).Find(&UserId)
	return UserId
}


func GainerLogin(user *model.LoginUser) int {
	cli := db.Get()
	var UserId int
	cli.Raw("select id from gainers where email= ? and passwd = ?", user.Email, user.Passwd).Find(&UserId)
	return UserId
}
