package v1

import (
	"A11Smile/deploy/db"
	"A11Smile/deploy/db/model"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"

	"strings"
)

func Register(user *model.User) error {
	key := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	passwd := user.Password

	//创建一个钱包用户
	create_account, err := key.NewAccount(passwd)
	if err != nil {
		return err
	}
	tracer := fmt.Sprintf("%v", create_account.URL)
	comma := strings.Index(tracer, "/")
	pos := strings.Index(tracer[comma:], "keystore")

	//完善用户钱包信息
	user.Key_store = tracer[comma+pos:]
	user.Block_address = fmt.Sprintf("%v", create_account.Address)

	// 引入数据库
	cli := db.Get()
	err = cli.Table("users").Save(user).Error
	if err != nil {
		return err
	}

	return nil
}
