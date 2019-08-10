package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type InfoController struct {
	controllers.BaseController
}

func (this *InfoController) Get() {
	this.Data["SelfArticleList"] = models.GetSelfArticleList(this.GetSession("UID").(int))

	this.TplName = "home/user/info.html"
}
