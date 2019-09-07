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
	c.Data["ArticleCount"] = models.ArticleCount(0, userID)
	c.Data["ArticleCollectCount"] = models.ArticleCollectCount(0, userID)

	c.TplName = "home/user/read.html"
}
