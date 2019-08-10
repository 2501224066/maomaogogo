package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type IndexController struct {
	controllers.BaseController
}

func (this *IndexController) Get() {
	this.Data["ArticleList"] = models.GetArticleList()

	this.TplName = "home/index.html"
}
