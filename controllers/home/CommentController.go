package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

// CommentController 评论控制器
type CommentController struct {
	controllers.BaseController
}

// 评论表单结构体
type commentForm struct {
	Content string `form:"editor" valid:"Required" chn:"回复内容"`
}

// Post ...
func (c *CommentController) Post() {
	var input commentForm
	c.ParseForm(&input)
	c.FormVerify(&input)

	articleID, _ := strconv.Atoi(c.Ctx.Input.Param(":article_id"))
	userID := c.GetSession("UID").(int)

	if b := models.CommentInsert(articleID, userID, input.Content); b == false {
		c.ResponseJSON(false, "发布失败")
	}

	// 文章评论数增加
	models.ArticleCommentNumUp(articleID)

	c.ResponseJSON(true, "发布成功")
}
