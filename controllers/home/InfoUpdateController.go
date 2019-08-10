package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type InfoUpdateController struct {
	controllers.BaseController
}

type InfoUpdateForm struct {
	Nickname  string `form:"nickname" valid:"Required;MaxSize(10)"`
	InShort   string `form:"in_short" valid:"Required;MaxSize(20)"`
	Sex       string `form:"sex" valid:"Required;MaxSize(1)"`
	Birthday  string `form:"birthday" valid:"Required;MaxSize(10)"`
	Introduce string `form:"introduce" valid:"Required;MaxSize(200)"`
}

func (this *InfoUpdateController) Get() {
	this.TplName = "home/user/save_info.html"
}

func (this *InfoUpdateController) Post() {
	var input InfoUpdateForm
	this.ParseForm(&input)
	this.FormVerify(&input)

	if re := models.UserUpdate(this.GetSession("UID").(int), input.Nickname, input.InShort, input.Sex, input.Birthday, input.Introduce); re == false {
		this.ResponseJson(false, "操作失败")
	}

	this.ResponseJson(true, "操作成功")
}
