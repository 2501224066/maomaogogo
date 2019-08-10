package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

type AvatarUrlUpdateController struct {
	controllers.BaseController
}

func (this *AvatarUrlUpdateController) Post() {
	path := this.GetString("path")

	re := models.AvatarUrlUpdate(this.GetSession("UID").(int), path)
	if re == false {
		this.ResponseJson(false, "操作失败")
	}

	this.ResponseJson(true, "操作成功")
}
