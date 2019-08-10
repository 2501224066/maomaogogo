package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type articleForm struct {
	Title   string `form:"title" valid:"Required;MaxSize(50)"`
	Content string `form:"editor" valid:"Required"`
	Node    string `form:"node" valid:"Required;MaxSize(10)"`
	Tag     string `form:"tag" valid:"Required"`
}

type CreateArticleController struct {
	controllers.BaseController
}

func (this *CreateArticleController) Get() {
	this.Data["Node"] = models.AllNode()
	this.TplName = "home/article/create_article.html"
}

func (this *CreateArticleController) Post() {
	var input articleForm
	this.ParseForm(&input)
	this.FormVerify(&input)

	if b := models.ArticleInsert(this.GetSession("UID").(int), input.Title, input.Content, input.Node, input.Tag); b == false {
		this.ResponseJson(false, "发布失败")
	}

	this.ResponseJson(true, "发布成功")
}
