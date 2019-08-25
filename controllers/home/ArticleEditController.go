package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

type ArticleEditController struct {
	controllers.BaseController
}

type articleEditForm struct {
	ArticleId int    `form:"article_id" valid:"Required" chn:"文章标记"`
	Title     string `form:"title" valid:"Required;MaxSize(50)" chn:"标题"`
	Content   string `form:"editor" valid:"Required" chn:"内容"`
	Node      string `form:"node" valid:"Required;MaxSize(10)" chn:"节点"`
	Tag       string `form:"tag" valid:"Required" chn:"标签"`
}

func (this *ArticleEditController) Get() {
	article_id, _ := strconv.Atoi(this.Ctx.Input.Param(":article_id"))
	article, err := models.ArticleRead(article_id)
	if err != nil {
		this.Ctx.Redirect(302, "/404")
	}

	this.Data["Node"] = models.AllNode()
	this.Data["Article"] = article
	this.TplName = "home/article/edit.html"
}

func (this *ArticleEditController) Post() {
	var input articleEditForm
	this.ParseForm(&input)
	this.FormVerify(&input)

	if b := models.ArticleUpdate(input.ArticleId, input.Title, input.Content, input.Node, input.Tag); b == false {
		this.ResponseJson(false, "保存失败")
	}

	this.ResponseJson(true, "编辑成功")

}
