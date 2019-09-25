package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

// AboutController 关于控制器
type AboutController struct {
	controllers.Controller
}

// Get ...
func (c *AboutController) Get() {
	key := c.Ctx.Input.Param(":key")
	SystemSetting := models.SystemSettingRead(key)

	c.Data["Name"] = SystemSetting.Name
	c.Data["Content"] = SystemSetting.Value
	c.TplName = "home/layout/about.html"
}
