package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type UserReadController struct {
	controllers.BaseController
}

func (this *UserReadController) Get() {
	this.Data["ArticleCount"] = models.ArticleCount(0, this.GetSession("UID").(int))
	this.Data["ArticleCollectCount"] = models.ArticleCollectCount(0, this.GetSession("UID").(int))

	this.TplName = "home/user/read.html"
}
