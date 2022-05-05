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


	err = eth.Init("0xE037b842989A0B0aD8A325Ed8Cd3B135733813AC")

	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
