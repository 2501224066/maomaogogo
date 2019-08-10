package home

import (
	"maomaogogo/controllers"
	"maomaogogo/helper"
	"maomaogogo/services"
	"time"
)

type SendEmailController struct {
	controllers.BaseController
}

func (this *SendEmailController) Post() {
	var (
		email             string = this.GetString("email")
		body              string
		randNum           string = helper.RandNum(6)
		codeKey           string = email + "_code"
		codeKeyCreateTime string = codeKey + "_create_time"
	)

	switch this.Ctx.Input.Param(":type") {
	case "register":
		body = "尊敬的用户：您的注册验证码为 " + randNum + " 。有效期为5分钟，请在有效期内使用。"
		this.SetSession(codeKey, randNum)
		this.SetSession(codeKeyCreateTime, time.Now().UnixNano()/1e9)
	}

	b := services.SendEmail(email, "MaomaoGogo_Email", body, "html")
	if !b {
		this.ResponseJson(false, "邮件发送失败，请稍后重试")
	}

	this.ResponseJson(true, "邮件已发送")
}
