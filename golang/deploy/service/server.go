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
		register.POST("/verifyBizlicense", gainer_register_verifyBizlicense)
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
		user.POST("/uploadMedical",user_uploadMedicalHandler)
		user.POST("/uploadReport",user_uploadReportHandler)
		user.POST("/readMedicalInformation", user_readMedicalInformation)
		user.POST("/editData",user_editDataHandler)
		user.POST("/editDataGetData",user_editDataGetDataHandler)
		user.POST("/editDataGetName",user_editDataGetNameHandler)
		user.POST("/userSeeData",user_SeeDataHandler)
	}

	// 征求者
	gainer := r.Group("/gainer")
	{
		gainer.POST("/")
		gainer.POST("/editGainerData",gainer_editDataHandler)
		gainer.POST("/gainerEditDataGetData",gainer_editDataGetDataHandler)
		gainer.POST("/gainerEditGetId",gainer_editDataGetIdHandler)
		gainer.POST("/gainerSeeData",gainer_SeeDataHandler)
	}

	err = r.Run(addr)
	return err
}
