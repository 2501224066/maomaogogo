package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

// UserOpController 用户操作控制器
type UserOpController struct {
	controllers.BaseController
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

// ArticleDel 删除文章
func (c *UserOpController) ArticleDel() {
	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))

	b := models.ArticleDel(articleID)
	if !b {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}
