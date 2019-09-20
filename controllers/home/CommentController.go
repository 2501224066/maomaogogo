package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"regexp"
	"strconv"
)

// CommentController 评论控制器
type CommentController struct {
	controllers.Controller
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

	// 向文章主人发通知
	article, _ := models.ArticleRead(articleID)
	user := models.UserRead(userID)
	models.AddNotice(article.User.UserID, "<a class='gray' href='/user_read/"+strconv.Itoa(user.UserID)+"'>"+user.Nickname+"</a>"+"在你的文章<a class='gray' href='/article_read/"+strconv.Itoa(article.ArticleID)+"'>"+article.Title+"</a>下发布了新评论")

	// 向@的人发送通知
	aUserIDArr := AnalysisCommentContent(input.Content)
	for _, v := range aUserIDArr {
		aUserID, _ := strconv.Atoi(v)
		models.AddNotice(aUserID, "<a class='gray' href='/user_read/"+strconv.Itoa(user.UserID)+"'>"+user.Nickname+"</a>"+"在文章<a class='gray' href='/article_read/"+strconv.Itoa(article.ArticleID)+"'>"+article.Title+"</a>的评论中 @ 了你")
	}

	c.ResponseJSON(true, "发布成功")
}

// AnalysisCommentContent 分析评论内容,匹配出@的人
func AnalysisCommentContent(content string) []string {
	re, _ := regexp.Compile(`<span>@<a href="/user_read/.*?"><input`)
	aUserIDArr := re.FindAllString(content, -1)

	for k, v := range aUserIDArr {
		re, _ := regexp.Compile(`<span>@<a href="/user_read/(.*)"><input`)
		aUserIDArr[k] = re.FindStringSubmatch(v)[1]
	}
	return aUserIDArr
}
