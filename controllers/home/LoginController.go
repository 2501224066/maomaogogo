package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

// LoginController 登录控制器
type LoginController struct {
	controllers.BaseController
}

// loginForm 登录表单结构体
type loginForm struct {
	Email    string `form:"email" valid:"Required;Email" chn:"邮箱"`
	Password string `form:"password" valid:"Required;MinSize(6);MaxSize(16)" chn:"密码"`
}

// Get ...
func (c *LoginController) Get() {
	c.TplName = "home/auth/login.html"
}

// Post ...
func (c *LoginController) Post() {
	var input loginForm
	c.ParseForm(&input)
	c.FormVerify(&input)

	b, uid := models.CheckLogin(input.Email, input.Password)
	if !b {
		c.ResponseJSON(false, "邮箱或密码错误")
	}

	c.SetSession("UID", uid)
	c.ResponseJSON(true, "登录成功")
}
