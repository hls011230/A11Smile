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


	err = eth.Init([]string{"0x8113A8a3EFEE4562354aBaA25168Ed58441970C5", "0x4bAFB7b1cE6a453Df78E7aD08A5aAad57B494365"})


	if err != nil {
		panic(err)
	}

	// 初始化web服务
	service.Start()
}
