package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

// articleCreateForm 文章创建表单结构体
type articleCreateForm struct {
	Title   string `form:"title" valid:"Required;MaxSize(50)" chn:"标题"`
	Content string `form:"editor" valid:"Required" chn:"内容"`
	NodeID  int    `form:"node_id" valid:"Required" chn:"节点"`
	Tag     string `form:"tag" valid:"Required" chn:"标签"`
}

// ArticleCreateController 文章控制器
type ArticleCreateController struct {
	controllers.Controller
}

// Get ...
func (c *ArticleCreateController) Get() {
	c.Data["Node"] = models.AllNode()
	c.TplName = "home/article/create.html"
}

// Post ...
func (c *ArticleCreateController) Post() {
	var input articleCreateForm
	c.ParseForm(&input)
	c.FormVerify(&input)

	if b := models.ArticleInsert(c.GetSession("UID").(int), input.NodeID, input.Title, input.Content, input.Tag); b == false {
		c.ResponseJSON(false, "发布失败")
	}

	c.ResponseJSON(true, "发布成功")
}
