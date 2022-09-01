package test

import (
	"net/smtp"
	"testing"

	"github.com/jordan-wright/email"
)

func TestMail(t *testing.T) {
	e := email.NewEmail()
	//From：谁发来的
	e.From = "Get <2683661364@qq.com>"
	//To：发给谁的
	e.To = []string{"2683661364@qq.com"}
	//主题，标题
	e.Subject = "验证码发送"
	//普通文本内容，支持html
	//e.Text = []byte("小朋友！！！")
	e.HTML = []byte("您的验证码:<b>123456</b>")
	//send方法：smtp.qq.com:587：QQ email相关的域名端口号 smtp.PlainAuth：第一个参数为空，第二个参数为自己的邮箱，第三个参数为授权码，下面有讲如何获取授权码

	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "2683661364@qq.com", "diagiibhciordiab", "smtp.qq.com"))

	//返回EOF可以关闭ssl重试
	// err := e.SendWithTLS("smtp.qq.com:587",
	// 	smtp.PlainAuth("", "2683661364@qq.com", "diagiibhciordiab", "smtp.qq.com"),
	// 	&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
		return
	}

}
