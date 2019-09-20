package home

import (
	"maomaogogo/controllers"
	"maomaogogo/helper"
	"maomaogogo/services"
	"time"
)

// SendEmailController 发送邮件控制器
type SendEmailController struct {
	controllers.Controller
}

// Post ...
func (c *SendEmailController) Post() {
	var (
		email             = c.GetString("email")
		body              string
		randNum           = helper.RandNum(6)
		codeKey           = email + "_code"
		codeKeyCreateTime = codeKey + "_create_time"
	)

	switch c.Ctx.Input.Param(":type") {
	case "register":
		body = "尊敬的用户：您的注册验证码为 " + randNum + " 。有效期为5分钟，请在有效期内使用。"
		c.SetSession(codeKey, randNum)
		c.SetSession(codeKeyCreateTime, time.Now().UnixNano()/1e9)
	}

	b := services.SendEmail(email, "MaomaoGogo_Email", body, "html")
	if !b {
		c.ResponseJSON(false, "邮件发送失败，请稍后重试")
	}

	c.ResponseJSON(true, "邮件已发送")
}
