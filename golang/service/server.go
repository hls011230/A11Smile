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

	}

	// 登录
	login := r.Group("/login")
	{
		login.POST("/user", user_loginHandler)
	}

	// 用户（分享者）
	user := r.Group("/user")
	{
		register = user.Group("/register")
		{
			register.POST("/", user_registerHandler)
		}

		user.POST("/verifyIDCard", user_verifyIDCardHandler)
		user.POST("/readMedicalInformation", user_readMedicalInformation)
<<<<<<< HEAD:golang/deploy/service/server.go
		user.POST("/editData",user_editDataHandler)
		user.POST("/editDataGetData",user_editDataGetDataHandler)
		user.POST("/editDataGetName",user_editDataGetNameHandler)
		user.POST("/userSeeData",user_SeeDataHandler)
=======
		user.POST("/uploadUser",user_uploadUser)
		user.POST("/UploadMedicalInformation",user_UploadMedicalInformation)
		user.POST("/CheckTheBalance",user_CheckTheBalance)
		user.POST("/CheckTheAS",user_CheckTheAS)

>>>>>>> e55850593e7864994334e03e96d839c0559e7b77:golang/service/server.go
	}

	// 征求者
	gainer := r.Group("/gainer")
	{
		register = gainer.Group("/register")
		{
			register.POST("/verifyBizlicense", gainer_register_verifyBizlicense)
			gainer.POST("/uploadGainer",gainer_uploadGainer)
			gainer.POST("/ReleaseMedicalInformation",gainer_ReleaseMedicalInformation)
			gainer.POST("/ReviewAndReward",gainer_ReviewAndReward)
			gainer.POST("/CheckTheBalance",gainer_CheckTheBalance)
			gainer.POST("/CheckTheAS",gainer_CheckTheAS)
		}

		gainer.POST("/")
		gainer.POST("/editGainerData",gainer_editDataHandler)
		gainer.POST("/gainerEditDataGetData",gainer_editDataGetDataHandler)
		gainer.POST("/gainerEditGetId",gainer_editDataGetIdHandler)
		gainer.POST("/gainerSeeData",gainer_SeeDataHandler)
	}

	err = r.Run(addr)
	return err
}
