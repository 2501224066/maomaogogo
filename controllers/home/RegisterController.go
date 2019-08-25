package home

import (
	"maomaogogo/controllers"
	"maomaogogo/helper"
	"maomaogogo/models"
	"time"

	"github.com/astaxie/beego"
)

type RegisterController struct {
	controllers.BaseController
}

type registerForm struct {
	Nickname       string `form:"nickname" valid:"Required;MaxSize(10)" chn:"昵称"`
	Email          string `form:"email" valid:"Required;Email" chn:"邮箱"`
	Password       string `form:"password" valid:"Required;MinSize(6);MaxSize(16)" chn:"密码"`
	RepeatPassword string `form:"repeat_password" valid:"Required;MinSize(6);MaxSize(16)" chn:"重复密码"`
	Code           string `form:"code" valid:"Required;Match(/[0-9]{6}/)" chn:"验证码"`
}

func (this *RegisterController) Get() {
	this.TplName = "home/auth/register.html"
}

func (this *RegisterController) Post() {
	var input registerForm
	this.ParseForm(&input)
	this.FormVerify(&input)

	var (
		codeKey           string = input.Email + "_code"
		codeKeyCreateTime string = codeKey + "_create_time"
	)
	codeForSession := this.GetSession(codeKey)
	codeCreateTimeForSession := this.GetSession(codeKeyCreateTime)
	defer this.DelSession(codeKey)
	defer this.DelSession(codeKeyCreateTime)

	if codeForSession == nil {
		this.ResponseJson(false, "验证码不存在")
	}

	codeLive, _ := beego.AppConfig.Int64("code_live")
	if (time.Now().UnixNano()/1e9 - codeCreateTimeForSession.(int64)) > codeLive {
		this.ResponseJson(false, "验证码过期")
	}

	if input.Code != codeForSession {
		this.ResponseJson(false, "验证码错误")
	}

	if input.Password != input.RepeatPassword {
		this.ResponseJson(false, "两次输入密码不一致")
	}

	if count := models.GetEmailCount(input.Email); count > 0 {
		this.ResponseJson(false, "邮箱已存在")
	}

	if b := models.UserInsert(input.Nickname, input.Email, helper.Md5(input.Password)); b == false {
		this.ResponseJson(false, "注册失败")
	}

	this.ResponseJson(true, "注册成功")
}
