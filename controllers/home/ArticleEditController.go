package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

// ArticleEditController 编辑文章控制器
type ArticleEditController struct {
	controllers.BaseController
}

// articleEditForm 表单结构体
type articleEditForm struct {
	ArticleID int    `form:"article_id" valid:"Required" chn:"文章标记"`
	Title     string `form:"title" valid:"Required;MaxSize(50)" chn:"标题"`
	Content   string `form:"editor" valid:"Required" chn:"内容"`
	Node      string `form:"node" valid:"Required;MaxSize(10)" chn:"节点"`
	Tag       string `form:"tag" valid:"Required" chn:"标签"`
}

// Get ...
func (c *ArticleEditController) Get() {
	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))
	article, err := models.ArticleRead(articleID)
	if err != nil {
		c.Ctx.Redirect(302, "/404")
	}

	c.Data["Node"] = models.AllNode()
	c.Data["Article"] = article
	c.TplName = "home/article/edit.html"
}

// Post ...
func (c *ArticleEditController) Post() {
	var input articleEditForm
	c.ParseForm(&input)
	c.FormVerify(&input)

	if b := models.ArticleUpdate(input.ArticleID, input.Title, input.Content, input.Node, input.Tag); b == false {
		c.ResponseJSON(false, "保存失败")
	}

	c.ResponseJSON(true, "编辑成功")
}
