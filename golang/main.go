package main

import (
	"A11Smile/db"
	"A11Smile/eth"
	"A11Smile/service"
)

// 程序的入口
func main() {

	// 初始化数据库
	err := db.Init()
	if err != nil {
		panic(err)
	}

	// 初始化区块链(传入合约地址)


	err = eth.Init("0xa18467dD5877E3932fe64c769C17122254f499Fd")

	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
