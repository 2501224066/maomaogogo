package home

import "maomaogogo/controllers"

// LogoutController 登出控制器
type LogoutController struct {
	controllers.Controller
}

// Get ...
func (c *LogoutController) Get() {
	c.DelSession("UID")
	c.Ctx.Redirect(302, "/")
}
