package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type LoginController struct {
	controllers.BaseController
}

type loginForm struct {
	Email    string `form:"email" valid:"Required;Email" chn:"邮箱"`
	Password string `form:"password" valid:"Required;MinSize(6);MaxSize(16)" chn:"密码"`
}

func (this *LoginController) Get() {
	this.TplName = "home/auth/login.html"
}

func (this *LoginController) Post() {
	var input loginForm
	this.ParseForm(&input)
	this.FormVerify(&input)

	b, uid := models.CheckLogin(input.Email, input.Password)
	if !b {
		this.ResponseJson(false, "邮箱或密码错误")
	}

	this.SetSession("UID", uid)
	this.ResponseJson(true, "登录成功")
}
