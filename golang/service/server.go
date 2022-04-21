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

	// 用户（分享者）
	user := r.Group("/user")
	{
		// 用户注册
		register = user.Group("/register")
		{
			register.POST("/", user_registerHandler)
		}

		// 用户登录
		login := user.Group("/login")
		{
			login.POST("/", user_loginHandler)
		}

		// 用户实名认证
		user.POST("/verifyIDCard", user_verifyIDCardHandler)

		// 查看用户Eth余额
		user.POST("/CheckTheBalance", user_CheckTheBalanceHandler)

		// 查看用户AS余额
		user.POST("/CheckTheAS", user_CheckTheAS)

		// 用户上传病历信息
		user.POST("/uploadMedicalHistory",user_uploadMedicalHistoryHandler)

		// 返回用户所有的病历信息
		user.POST("/viewAllMedicalHistory",user_viewAllMedicalHistory)

		// 用户上传体检报告
		user.POST("/uploadMedicalExaminationReport",user_uploadMedicalExaminationReportHandler)

		// 返回用户所有的病历信息
		user.POST("/viewAllMedicalExaminationReport",user_viewAllMedicalExaminationReport)

		// 返回用户所有的病历信息

		user.POST("/readMedicalInformation", user_readMedicalInformation)
		user.POST("/editUserName", user_editUserNameHandler)
		user.POST("/editUserResume",user_editUserResumeHandler)
		user.POST("/editDataGetData", user_editDataGetDataHandler)
		user.POST("/editDataGetId", user_editDataGetIdHandler)
		user.POST("/userSeeData", user_SeeDataHandler)
		user.POST("/User_AuthenticationSee", User_AuthenticationSeeHandler)
		user.POST("/UploadMedicalInformation", user_UploadMedicalInformation)



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
		gainer.POST("/gainerSeeData", gainer_SeeDataHandler)
	}

	r.Run(":8080")
}
