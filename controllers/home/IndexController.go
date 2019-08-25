package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type IndexController struct {
	controllers.BaseController
}

func (this *IndexController) Get() {
	this.Data["ArticleList"] = models.GetArticleList(0, 0, 1)
	this.TplName = "home/index.html"
}
