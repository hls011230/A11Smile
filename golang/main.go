package main

import (
	"A11Smile/deploy/cfg"
	"A11Smile/deploy/db"
	"A11Smile/deploy/service"
	"fmt"
	"os"
)

// 程序的入口
func main() {

	// 初始化数据库
	err := db.Init()
	if err != nil {
		panic(err)
	}

	// 初始化配置文件
	c, err := cfg.LoadConfig(os.Args[1])
	if err != nil {
		fmt.Printf("载入配置文件错误:%v\n", err)
		return
	}

	// 初始化web服务
	service.Start(fmt.Sprintf("%s:%s", c.Host, c.Port), c.ContractAddress)
}
