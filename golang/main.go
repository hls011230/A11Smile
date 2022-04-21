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
	err = eth.Init("0xea98D3a0F722bE026Ac78ae45E6f433A9e6955FC")
	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
