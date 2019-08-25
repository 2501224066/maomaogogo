package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type articleCreateForm struct {
	Title   string `form:"title" valid:"Required;MaxSize(50)" chn:"标题"`
	Content string `form:"editor" valid:"Required" chn:"内容"`
	Node    string `form:"node" valid:"Required;MaxSize(10)" chn:"节点"`
	Tag     string `form:"tag" valid:"Required" chn:"标签"`
}

type ArticleCreateController struct {
	controllers.BaseController
}

func (this *ArticleCreateController) Get() {
	this.Data["Node"] = models.AllNode()
	this.TplName = "home/article/create.html"
}

func (this *ArticleCreateController) Post() {
	var input articleCreateForm
	this.ParseForm(&input)
	this.FormVerify(&input)

	if b := models.ArticleInsert(this.GetSession("UID").(int), input.Title, input.Content, input.Node, input.Tag); b == false {
		this.ResponseJson(false, "发布失败")
	}

	this.ResponseJson(true, "发布成功")
}
