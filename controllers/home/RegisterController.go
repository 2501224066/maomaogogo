package home

import (
	"maomaogogo/controllers"
	"maomaogogo/helper"
	"maomaogogo/models"
	"time"

	"github.com/astaxie/beego"
)

// RegisterController 注册控制器
type RegisterController struct {
	controllers.Controller
}

// registerForm 注册表单结构体
type registerForm struct {
	Nickname       string `form:"nickname" valid:"Required;MaxSize(10)" chn:"昵称"`
	Email          string `form:"email" valid:"Required;Email" chn:"邮箱"`
	Password       string `form:"password" valid:"Required;MinSize(6);MaxSize(16)" chn:"密码"`
	RepeatPassword string `form:"repeat_password" valid:"Required;MinSize(6);MaxSize(16)" chn:"重复密码"`
	Code           string `form:"code" valid:"Required;Match(/[0-9]{6}/)" chn:"验证码"`
}

// Get ...
func (c *RegisterController) Get() {
	c.TplName = "home/auth/register.html"
}

// Post ...
func (c *RegisterController) Post() {
	var input registerForm
	c.ParseForm(&input)
	c.FormVerify(&input)

	var (
		codeKey           = input.Email + "_code"
		codeKeyCreateTime = codeKey + "_create_time"
	)
	codeForSession := c.GetSession(codeKey)
	codeCreateTimeForSession := c.GetSession(codeKeyCreateTime)
	defer c.DelSession(codeKey)
	defer c.DelSession(codeKeyCreateTime)

	if codeForSession == nil {
		c.ResponseJSON(false, "验证码不存在")
	}

	codeLive, _ := beego.AppConfig.Int64("code_live")
	if (time.Now().UnixNano()/1e9 - codeCreateTimeForSession.(int64)) > codeLive {
		c.ResponseJSON(false, "验证码过期")
	}

	if input.Code != codeForSession {
		c.ResponseJSON(false, "验证码错误")
	}

	if input.Password != input.RepeatPassword {
		c.ResponseJSON(false, "两次输入密码不一致")
	}

	if count := models.GetEmailCount(input.Email); count > 0 {
		c.ResponseJSON(false, "邮箱已存在")
	}

	if b := models.UserInsert(input.Nickname, input.Email, helper.Md5(input.Password)); b == false {
		c.ResponseJSON(false, "注册失败")
	}

	c.ResponseJSON(true, "注册成功")
}
