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
	err = eth.Init("0x07FFDc22cb928f3dadfa5B8473F3A841A1E4E8aC")
	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
