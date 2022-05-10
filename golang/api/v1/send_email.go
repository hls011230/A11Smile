package v1

import (
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

// 发送邮件
func SendMail(mailTo string,subject string, body string ) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string {
		"user": "3026683545@qq.com",
		"pass": "yeeipzolhsjpdcdi",
		"host": "smtp.qq.com",
		"port": "465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From","A11Smile.abi" + "<" + mailConn["user"] + ">")  //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo)  //发送给多个用户
	m.SetHeader("Subject", subject)  //设置邮件主题
	m.SetBody("text/html", body)     //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}

// 随机生成6位数验证码
func Create_code() string {
	rand.Seed(time.Now().Unix())
	code := ""
	for i:= 0; i < 6;i ++ {
		num := rand.Intn(10)
		str1 := strconv.Itoa(num)
		code += str1
	}
	return code
}


