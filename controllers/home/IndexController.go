package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

// IndexController 主页控制器
type IndexController struct {
	controllers.BaseController
}

// Get ...
func (c *IndexController) Get() {
	c.Data["ArticleList"] = models.ArticleList(0, 0, 1)

	if userID := c.GetSession("UID"); userID != nil {
		c.Data["UnreadNoticeCount"] = models.UnreadNoticeCount(userID.(int), 1)
	}
	c.TplName = "home/index.html"
}
