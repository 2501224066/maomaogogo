package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

// UserResourceController 用户资源控制器
type UserResourceController struct {
	controllers.BaseController
}

// Article 用户文章
func (c *UserResourceController) Article() {
	// 当前页面
	p, _ := strconv.Atoi(c.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":user_id"))

	c.Data["UserInfo"] = models.UserRead(userID)
	c.Data["PageNo"] = p
	c.Data["Count"] = models.ArticleCount(0, userID)
	c.Data["ArticleList"] = models.ArticleList(0, userID, p)
	c.TplName = "home/user/article.html"
}
