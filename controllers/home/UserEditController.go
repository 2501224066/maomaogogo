package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type UserEditController struct {
	controllers.BaseController
}

type userEditForm struct {
	Nickname  string `form:"nickname" valid:"Required;MaxSize(10)" chn:"昵称"`
	InShort   string `form:"in_short" valid:"Required;MaxSize(20)" chn:"简介"`
	Sex       string `form:"sex" valid:"Required;MaxSize(1)" chn:"性别"`
	Birthday  string `form:"birthday" valid:"Required;MaxSize(10)" chn:"生日"`
	Introduce string `form:"introduce" valid:"Required;MaxSize(200)" chn:"详细介绍"`
}

func (this *UserEditController) Get() {
	this.TplName = "home/user/edit.html"
}

func (this *UserEditController) Post() {
	var input userEditForm
	this.ParseForm(&input)
	this.FormVerify(&input)

	if re := models.UserUpdate(this.GetSession("UID").(int), input.Nickname, input.InShort, input.Sex, input.Birthday, input.Introduce); re == false {
		this.ResponseJson(false, "操作失败")
	}

	this.ResponseJson(true, "操作成功")
}

func (this *UserEditController) AvatarUrlUpdata() {
	path := this.GetString("path")

	re := models.AvatarUrlUpdate(this.GetSession("UID").(int), path)
	if re == false {
		this.ResponseJson(false, "操作失败")
	}

	this.ResponseJson(true, "操作成功")
}

func (this *UserEditController) QrImgUpdata() {
	path := this.GetString("path")

	re := models.QrImgUpdate(this.GetSession("UID").(int), path)
	if re == false {
		this.ResponseJson(false, "操作失败")
	}

	this.ResponseJson(true, "操作成功")
}
