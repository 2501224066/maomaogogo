package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
	"strconv"
)

type UserOpController struct {
	controllers.BaseController
}

func (this *UserOpController) ArticleLike() {
	if u := this.GetSession("UID"); u == nil {
		this.ResponseJson(false, "未登录禁止操作")
	}

	article_id, _ := strconv.Atoi(this.Ctx.Input.Param(":article_id"))
	uid := this.GetSession("UID").(int)

	// 点赞状态
	count := models.ArticleLikeCount(article_id, uid)

	var b bool
	if count > 0 {
		b = models.ArticleLikeDown(article_id, uid)
	} else {
		b = models.ArticleLikeUp(article_id, uid)
	}

	if !b {
		this.ResponseJson(false, "操作失败")
	}

	this.ResponseJson(true, "操作成功")
}
