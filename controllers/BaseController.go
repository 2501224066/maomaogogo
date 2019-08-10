package controllers

import (
	"html/template"
	"maomaogogo/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type BaseController struct {
	beego.Controller
}

var (
	LoginStatus bool
	User        models.User
)

func (this *BaseController) Prepare() {
	// 路径
	this.Data["Path"] = this.Ctx.Request.RequestURI

	// CSRF
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	// 登录状态
	uid := this.GetSession("UID")
	if uid != nil {
		LoginStatus = true
		User = models.UserRead(uid.(int))
	} else {
		LoginStatus = false
	}
	this.Data["LoginStatus"] = LoginStatus
	this.Data["User"] = User

	this.Layout = "home/layout/layout.html"
}

// 返回JSON
func (this *BaseController) ResponseJson(isSuccess bool, msg string, data ...interface{}) {
	status := 0
	if isSuccess {
		status = 1
	}
	ret := map[string]interface{}{"status": status, "msg": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	this.Data["json"] = ret
	this.ServeJSON()
	this.StopRun()
}

// 表单校验
func (this *BaseController) FormVerify(input interface{}) {
	valid := validation.Validation{}
	b, _ := valid.Valid(input)
	if !b {
		this.ResponseJson(false, valid.Errors[0].Key+valid.Errors[0].Message)
	}
}
