package home

import (
	"maomaogogo/controllers"
)

type LogoutController struct {
	controllers.BaseController
}

func (this *LogoutController) Get() {
	this.DelSession("UID")
	this.Ctx.Redirect(302, "/")
}
