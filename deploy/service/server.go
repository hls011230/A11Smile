package service

import "github.com/gin-gonic/gin"

// 合约地址
var Contract_address string

// 启动服务
func Start(addr,contract_address string) (err error) {
	r := gin.Default()

	// 初始化合约地址
	Contract_address = contract_address


	register := r.Group("/register")
	{
		register.POST("/sendEmail",user_register_sendEmailHandler)
		register.POST("/verifyEmail",user_register_verifyEmailHandler)
		register.POST("/",user_registerHandler)
	}


	err = r.Run(addr)
	return err
}