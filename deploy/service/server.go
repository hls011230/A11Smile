package service

import "github.com/gin-gonic/gin"

// 合约地址
var Contract_address string

// 启动服务
func Start(addr, contract_address string) (err error) {
	r := gin.Default()

	// 初始化合约地址
	Contract_address = contract_address

	// 初始化静态文件夹
	r.Static("static", "../static")

	// 注册
	register := r.Group("/register")
	{
		register.POST("/sendEmail", user_register_sendEmailHandler)
		register.POST("/verifyEmail", user_register_verifyEmailHandler)
		register.POST("/user", user_registerHandler)
	}

	// 登录
	login := r.Group("/login")
	{
		login.POST("/user", user_loginHandler)
	}

	// 用户（分享者）
	user := r.Group("/user")
	{
		user.POST("/verifyIDCard", user_verifyIDCardHandler)
	}

	// 征求者
	gainer := r.Group("/gainer")
	{
		gainer.POST("/")
	}

	err = r.Run(addr)
	return err
}
