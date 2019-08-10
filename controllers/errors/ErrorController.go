package errors

import "maomaogogo/controllers"

type ErrorController struct {
	controllers.BaseController
}

func (this *ErrorController) Error404() {
	this.Data["Content"] = "啊！怎么不见了..."
	this.TplName = "error/404.html"
}
