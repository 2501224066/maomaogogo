package services

import (
	"net/smtp"
	"strings"

	"github.com/astaxie/beego"
)

/**
 * 发送邮件
 * @to string 目标邮箱
 * @subject string 邮件标题
 * @body string 邮件内容
 * @mailtype string 邮件格式
 */
func SendEmail(to, subject, body, mailtype string) bool {
	var user = beego.AppConfig.String("email::email_user")
	var password = beego.AppConfig.String("email::email_password")
	var host = beego.AppConfig.String("email::email_host")

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	if err != nil {
		return false
	}
	return true
}
