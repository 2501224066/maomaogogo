package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

type ReadArticleController struct {
	controllers.BaseController
}

func (this *ReadArticleController) Get() {
	// 文章信息
	id_str := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(id_str)
	article, err := models.ArticleRead(id)
	if err != nil {
		this.Ctx.Redirect(302, "/")
	}

	// 文章阅读 +1
	models.ArticleReadAddOne(article)

	this.Data["Article"] = article
	this.TplName = "home/article/read_article.html"
}
