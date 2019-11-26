package utils

import (
	"fmt"
	"goAdmin/src/main/comm/config"
	"gopkg.in/gomail.v2"
	"strings"
)

func SendMail(mailTo string, subject string, body string) error {
	notifyOpen := config.CommConfig.Notify
	if !notifyOpen.ErrorIsOpen {
		return nil
	}
	emailMessage := gomail.NewMessage()
	emailConfig := config.CommConfig.Notify.EmailConfig
	//设置发件人
	emailMessage.SetHeader("From", emailConfig.FromEmail)
	//设置发送给多个用户
	mailArrTo := strings.Split(mailTo, ",")
	emailMessage.SetHeader("To", mailArrTo...)
	//设置邮件主题
	emailMessage.SetHeader("Subject", subject)
	//设置邮件正文
	emailMessage.SetBody("text/html", body)

	d := gomail.NewDialer(emailConfig.EmailHost, emailConfig.EmailPort, emailConfig.FromEmail, emailConfig.EmailPass)
	err := d.DialAndSend(emailMessage)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
