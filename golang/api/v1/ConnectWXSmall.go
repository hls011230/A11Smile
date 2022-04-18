package v1

import (
	"A11Smile/db/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetToken() (model.RespWXToken, error) {
	// Get请求获取接口token
	var token model.RespWXToken
	url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	res, err := http.DefaultClient.Get(fmt.Sprintf(url, model.WXAPP.AppId, model.WXAPP.Secret))
	if err != nil {
		return token, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {
		return token, err
	}
	return token, err
}
