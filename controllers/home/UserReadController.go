package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type UserReadController struct {
	controllers.BaseController
}

func (this *UserReadController) Get() {
	this.Data["SelfArticleList"] = models.GetArticleList(0, this.GetSession("UID").(int), 1)
	this.TplName = "home/user/read.html"
}
