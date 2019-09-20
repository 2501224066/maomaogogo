package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"

	"strconv"
)

// ArticleReadController 阅读文章控制器
type ArticleReadController struct {
	controllers.Controller
}

// Get ...
func (c *ArticleReadController) Get() {
	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))
	article, err := models.ArticleRead(articleID)
	if err != nil {
		c.Ctx.Redirect(302, "/404")
	}

	// 阅读数 +1
	models.ArticleReadNumUp(article)

	u := c.GetSession("UID")
	// 点赞状态
	c.Data["LikeStatus"] = false
	if u != nil {
		count := models.ArticleLikeCount(articleID, c.GetSession("UID").(int))
		if count > 0 {
			c.Data["LikeStatus"] = true
		}
	}

	// 收藏状态
	c.Data["CollectStatus"] = false
	if u != nil {
		count := models.ArticleCollectCount(articleID, c.GetSession("UID").(int))
		if count > 0 {
			c.Data["CollectStatus"] = true
		}
	}

	// 评论
	c.Data["CommentList"] = models.CommentList(articleID)

	c.Data["Article"] = article
	c.TplName = "home/article/read.html"
}
