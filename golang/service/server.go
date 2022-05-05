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
		user.POST("/viewAllMedicalHistory",user_viewAllMedicalHistoryHandler)

		// 用户预览病历信息
		user.POST("/previewMedicalHistory",user_previewMedicalHistoryHandler)

		// 用户上传体检报告
		user.POST("/uploadMedicalExaminationReport",user_uploadMedicalExaminationReportHandler)

		// 返回用户所有的体检报告
		user.POST("/viewAllMedicalExaminationReport",user_viewAllMedicalExaminationReportHandler)

		// 用户预览体检报告信息
		user.POST("/previewMedicalExaminationReport",user_previewMedicalExaminationReportHandler)

		// 展示用户的个人信息
		user.POST("/userAuthenticationSee", user_authenticationSeeHandler)

		// 修改用户名
		user.POST("/editUserName", user_editUserNameHandler)

		// 修改用户个人简介
		user.POST("/editUserResume",user_editUserResumeHandler)

		user.POST("/readMedicalInformation", user_readMedicalInformation)
		user.POST("/UploadMedicalInformation", user_UploadMedicalInformation)

		// 用户首页展示
		user.POST("/userDisplayHomepage",user_displayHomepageHandler)

		// 用户详情界面展示
		user.POST("/showDetailsPage",user_showDetailsPageHandler)

		//用户上传证书
		user.POST("/SubmitCertificate",user_submitCertificateHandler)

		//用户已经完成交易展示
		user.POST("/AllTransactions",user_showAllTransactionsHandler)

		//用户未交易完成展示
		user.POST("/NoTransactions",user_showNoTransactionsHandler)

	}

	// 征求者
	gainer := r.Group("/gainer")
	{
		// 征求者注册
		register = gainer.Group("/register")
		{
			// 营业执照认证
			register.POST("/verifyBizlicense", gainer_register_verifyBizlicenseHandler)

			register.POST("/",gainer_registerHandler)
		}

		// 征求者登录
		gainer.POST("/login",gainer_loginHandler)

		//gainer.POST("/uploadGainer", gainer_uploadGainer)

		// 征求者发布征求信息
		gainer.POST("/ReleaseMedicalInformation", gainer_ReleaseMedicalInformation)


		//gainer.POST("/ReviewAndReward", gainer_ReviewAndReward)

		// 查看征求者ETH余额
		gainer.POST("/CheckTheBalance", gainer_CheckTheBalance)

		// 查看征求者AS余额
		gainer.POST("/CheckTheAS", gainer_CheckTheAS)

		// 征求者编辑个人资料
		gainer.POST("/gainerEdit", gainer_editDataHandler)

		// 征求者个人资料展示
		gainer.POST("/gainerAauthenticationSee", gainer_authenticationSeeHandler)

		// 征求者首页展示
		gainer.POST("/gainerDisplayHomepage",gainer_displayHomepageHandler)

		// 征求者头像设置
		gainer.POST("/editGainerIcon",gainer_editGainerIconHandler)

		// 征求者头像展示
		gainer.POST("/showGainerIcon",gainer_showGainerIconHandler)

		//征求者审核功能
		gainer.POST("/Examine",gainer_ExamineHander)

		//征求者查看用户上传的证书
		gainer.POST("/ViewCertificate",gainer_ViewCertificateHandler)


	}

	r.Run(":8080")
}
