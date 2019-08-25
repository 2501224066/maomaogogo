package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

type ArticleReadController struct {
	controllers.BaseController
}

func (this *ArticleReadController) Get() {
	article_id, _ := strconv.Atoi(this.Ctx.Input.Param(":article_id"))
	article, err := models.ArticleRead(article_id)
	if err != nil {
		this.Ctx.Redirect(302, "/404")
	}

	// 阅读数 +1
	models.ArticleReadAddOne(article)

	// 点赞状态
	this.Data["LikeStatus"] = false
	u := this.GetSession("UID")
	if u != nil {
		count := models.ArticleLikeCount(article_id, this.GetSession("UID").(int))
		if count > 0 {
			this.Data["LikeStatus"] = true
		}
	}

	this.Data["Article"] = article
	this.TplName = "home/article/read.html"
}
