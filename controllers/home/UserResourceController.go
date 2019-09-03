package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

type UserResourceController struct {
	controllers.BaseController
}

func (this *UserResourceController) Article() {
	// 当前页面
	p, _ := strconv.Atoi(this.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	this.Data["PageNo"] = p
	this.Data["Count"] = models.ArticleCount(0, this.GetSession("UID").(int))
	this.Data["ArticleList"] = models.ArticleList(0, this.GetSession("UID").(int), p)
	this.TplName = "home/user/article.html"
}
