package service

import (
	v1 "A11Smile/deploy/api/v1"
	"A11Smile/deploy/db/model"
	"A11Smile/deploy/serializer"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 用户注册
func user_registerHandler(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		serializer.RespError(c, fmt.Sprintf("数据绑定错误:%v", err))
		return
	}

	err := v1.Register(&user)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "用户注册成功")

}

// 邮箱对应验证码
var Code = make(map[string]string)

// 用户发送邮箱验证码
func user_register_sendEmailHandler(c *gin.Context) {
	// 获取用户发送的邮箱
	mailTo := c.PostForm("email")

	// 随机生成六位数验证码
	Code[mailTo] = v1.Create_code()
	// 邮件正文
	body := fmt.Sprintf(`
				<html>
				<body>
				<h3>
				"你的验证码是：%v"
				</h3>
				</body>
				</html>
	
			`, Code[mailTo])
	err := v1.SendMail(mailTo, "用户注册", body)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "邮箱发送成功！")
}

// 用户验证邮箱验证码
func user_register_verifyEmailHandler(c *gin.Context) {

	// 获取用户发送的邮箱
	mailTo := c.PostForm("email")
	code := c.PostForm("code")

	if code != Code[mailTo] {
		serializer.RespError(c, "验证码错误")
		return
	}

	// 验证成功后删除映射中所对应的账号
	delete(Code, mailTo)
	serializer.RespOK(c, "验证码正确")
}
