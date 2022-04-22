package service

import "github.com/gin-gonic/gin"



// 启动服务
func Start()  {
	r := gin.Default()

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

		user.POST("/editUserName", user_editUserNameHandler)
		user.POST("/editUserResume",user_editUserResumeHandler)
		user.POST("/userAuthenticationSee", user_authenticationSeeHandler)

		user.POST("/UploadMedicalInformation", user_UploadMedicalInformation)
		user.POST("/CheckTheBalance", user_CheckTheBalanceHandler)
		user.POST("/CheckTheAS", user_CheckTheAS)

		user.POST("/uploadMedicalHistory",user_uploadMedicalHistoryHandler)


	}

	// 征求者
	gainer := r.Group("/gainer")
	{
		register = gainer.Group("/register")
		{
			register.POST("/verifyBizlicense", gainer_register_verifyBizlicense)

		}

		gainer.POST("/uploadGainer", gainer_uploadGainer)
		gainer.POST("/ReleaseMedicalInformation", gainer_ReleaseMedicalInformation)
		gainer.POST("/ReviewAndReward", gainer_ReviewAndReward)
		gainer.POST("/CheckTheBalance", gainer_CheckTheBalance)
		gainer.POST("/CheckTheAS", gainer_CheckTheAS)
		gainer.POST("/")
		gainer.POST("/editGainerData", gainer_editDataHandler)
		gainer.POST("/gainerEditDataGetData", gainer_editDataGetDataHandler)
		gainer.POST("/gainerEditGetId", gainer_editDataGetIdHandler)
		gainer.POST("/gainerAauthenticationSee", gainer_authenticationSeeHandler)
	}

	r.Run(":8080")

}
