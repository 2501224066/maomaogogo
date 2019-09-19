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

// CollectArticle 收藏的文章
func (c *UserResourceController) CollectArticle() {
	// 当前页面
	p, _ := strconv.Atoi(c.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":user_id"))

	userInfo := models.UserRead(userID)
	c.Data["UserInfo"] = userInfo
	c.Data["PageNo"] = p
	c.Data["Count"] = models.ArticleCollectCount(0, userID)
	c.Data["CollectArticleList"] = models.CollectArticleList(userID, p)
	c.TplName = "home/user/collect_article.html"
}

// FollowUser 关注的用户
func (c *UserResourceController) FollowUser() {
	// 当前页面
	p, _ := strconv.Atoi(c.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":user_id"))

	userInfo := models.UserRead(userID)
	c.Data["UserInfo"] = userInfo
	c.Data["PageNo"] = p
	c.Data["Count"] = userInfo.FollowNum
	c.Data["FollowUserList"] = models.FollowUserList(userID, p)
	c.TplName = "home/user/follow_user.html"
}

// Fans 粉丝
func (c *UserResourceController) Fans() {
	// 当前页面
	p, _ := strconv.Atoi(c.Input().Get("p"))
	if p == 0 {
		p = 1
	}

	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":user_id"))

	userInfo := models.UserRead(userID)
	c.Data["UserInfo"] = userInfo
	c.Data["PageNo"] = p
	c.Data["Count"] = userInfo.FansNum
	c.Data["FansList"] = models.FansList(userID, p)
	c.TplName = "home/user/fans.html"
}
