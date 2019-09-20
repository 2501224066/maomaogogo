package errors

import "maomaogogo/controllers"

// ErrorController 错误控制器
type ErrorController struct {
	controllers.Controller
}

// Error404 404错误
func (c *ErrorController) Error404() {
	c.Data["Content"] = "啊！怎么不见了..."
	c.TplName = "error/404.html"
}
