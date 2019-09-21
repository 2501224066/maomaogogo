package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

// ArticleListController 文章列表结构体
type ArticleListController struct {
	controllers.Controller
}

// Get ...
func (c *ArticleListController) Get() {
	// 当前节点
	nodeID, _ := strconv.Atoi(c.Ctx.Input.Param(":node_id"))

	// 当前页面
	p, _ := strconv.Atoi(c.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	c.Data["PageNo"] = p
	c.Data["Count"] = models.ArticleCount(nodeID, 0)
	c.Data["ArticleList"] = models.ArticleList(nodeID, 0, p)
	c.Data["Node"] = models.AllNode()
	c.Data["NowNodeId"] = nodeID
	c.TplName = "home/article/list.html"
}

// IndexPage 网站入口
func (c *ArticleListController) IndexPage() {
	c.Data["ArticleList"] = models.ArticleList(0, 0, 1)
	c.TplName = "home/index.html"
}

// SearchArticle 搜索文章
func (c *ArticleListController) SearchArticle() {
	// 关键词
	keyword := c.Ctx.Input.Param(":keyword")

	// 当前页面
	p, _ := strconv.Atoi(c.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	c.Data["PageNo"] = p
	c.Data["Count"] = models.SearchArticleCount(keyword)
	c.Data["ArticleList"] = models.SearchArticleList(keyword, p)
	c.Data["Keyword"] = keyword
	c.TplName = "home/article/search.html"
}
