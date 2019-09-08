package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

// UserReadController 查看用户控制器
type UserReadController struct {
	controllers.BaseController
}

// Get ...
func (c *UserReadController) Get() {
	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":user_id"))

	c.Data["UserInfo"] = models.UserRead(userID)
	// 文章数
	c.Data["ArticleCount"] = models.ArticleCount(0, userID)
	// 收藏文章数
	c.Data["ArticleCollectCount"] = models.ArticleCollectCount(0, userID)
	// 用户所有文章数据
	c.Data["AllArticleDataCount"] = models.AllArticleDataCount(userID)

	// 关注状态
	c.Data["FollowStatus"] = false
	count := models.FollowCount(c.GetSession("UID").(int), userID)
	if count > 0 {
		c.Data["FollowStatus"] = true
	}

	c.TplName = "home/user/read.html"
}
