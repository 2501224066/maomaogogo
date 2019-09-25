package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

// UserOpController 用户操作控制器
type UserOpController struct {
	controllers.Controller
}

// ArticleLike 文章点赞
func (c *UserOpController) ArticleLike() {
	if u := c.GetSession("UID"); u == nil {
		c.ResponseJSON(false, "未登录禁止操作")
	}

	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))
	userID := c.GetSession("UID").(int)

	// 点赞状态
	count := models.ArticleLikeCount(articleID, userID)

	var b bool
	if count > 0 {
		b = models.ArticleLikeDown(articleID, userID)
	} else {
		b = models.ArticleLikeUp(articleID, userID)
	}

	if !b {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}

// ArticleCollect 文章收藏
func (c *UserOpController) ArticleCollect() {
	if u := c.GetSession("UID"); u == nil {
		c.ResponseJSON(false, "未登录禁止操作")
	}

	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))
	userID := c.GetSession("UID").(int)

	// 收藏状态
	count := models.ArticleCollectCount(articleID, userID)

	var b bool
	if count > 0 {
		b = models.ArticleCollectDown(articleID, userID)
	} else {
		b = models.ArticleCollectUp(articleID, userID)
	}

	if !b {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}

// ArticleReport 文章举报
func (c *UserOpController) ArticleReport() {
	if u := c.GetSession("UID"); u == nil {
		c.ResponseJSON(false, "未登录禁止操作")
	}

	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))
	reportNum := models.ArticleReportUp(articleID)
	models.ArticleReportBan(reportNum, articleID)

	c.ResponseJSON(true, "已接到您的举报，我们会尽快处理！")
}

// ArticleDel 删除文章
func (c *UserOpController) ArticleDel() {
	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))

	b := models.ArticleDel(articleID)
	if !b {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}

// Follow 用户关注
func (c *UserOpController) Follow() {
	userID, _ := strconv.Atoi(c.Ctx.Input.Param(":user_id"))
	if userID == 0 {
		c.ResponseJSON(false, "操作失败")
	}

	// 关注状态
	count := models.FollowCount(c.GetSession("UID").(int), userID)

	var b bool
	if count > 0 {
		b = models.UserFollowDown(c.GetSession("UID").(int), userID)
	} else {
		b = models.UserFollowUp(c.GetSession("UID").(int), userID)
	}

	if !b {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}
