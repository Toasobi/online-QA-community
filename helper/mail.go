package helper

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	//From：谁发来的
	e.From = "Get <2683661364@qq.com>"
	//To：发给谁的
	e.To = []string{toUserEmail}
	//主题，标题
	e.Subject = "用户注册消息"

	e.HTML = []byte("您的验证码:<b>" + code + "</b>")

	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "2683661364@qq.com", "diagiibhciordiab", "smtp.qq.com"))

	if err != nil {
		return err
	}
	return nil
}
