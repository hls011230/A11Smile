package v1

import (
	"A11Smile/deploy/db"
	"A11Smile/deploy/db/model"
	"fmt"

	"net/http"
)

func Login(user *model.LoginUser) int {
	cli := db.Get()
	var count int
	cli.Raw("select count(1) from users where email= ? and password = ?", user.Email, user.Password).Find(&count)
	return count
}

func LoginWXSmall(code string) (wxInfo interface{}, err error) {
	appId := "wx1f6551a9b737be68"
	appSecret := "e8c9ec131ccbcc684b6cd0327404c869"
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	resp, err := http.Get(fmt.Sprintf(url, appId, appSecret, code))
	if err != nil {
		return wxInfo, err
	}
	defer resp.Body.Close()

	if err != nil {
		return wxInfo, err
	}

	wxInfo = resp
	return wxInfo, nil
}
