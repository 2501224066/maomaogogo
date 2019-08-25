package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

type ArticleListController struct {
	controllers.BaseController
}

func (this *ArticleListController) Get() {
	// 当前节点
	node_id_str := this.Ctx.Input.Param(":node_id")
	node_id, _ := strconv.Atoi(node_id_str)

	// 当前页面
	p, _ := strconv.Atoi(this.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	this.Data["PageNo"] = p
	this.Data["Count"] = models.GetArticleCount(node_id, 0)
	this.Data["ArticleList"] = models.GetArticleList(node_id, 0, p)
	this.Data["Node"] = models.AllNode()
	this.Data["NowNodeId"] = node_id
	this.TplName = "home/article/list.html"
}
