package home

import (
	"maomaogogo/controllers"
	"maomaogogo/models"
)

// UserEditController 编辑用户控制器
type UserEditController struct {
	controllers.Controller
}

// userEditForm 编辑用户表单结构体
type userEditForm struct {
	Nickname  string `form:"nickname" valid:"Required;MaxSize(10)" chn:"昵称"`
	InShort   string `form:"in_short" valid:"Required;MaxSize(20)" chn:"简介"`
	Sex       string `form:"sex" valid:"Required;MaxSize(1)" chn:"性别"`
	Birthday  string `form:"birthday" valid:"Required;MaxSize(10)" chn:"生日"`
	Introduce string `form:"introduce" valid:"Required;MaxSize(200)" chn:"详细介绍"`
}

// Get ...
func (c *UserEditController) Get() {
	c.TplName = "home/user/edit.html"
}

// Post ...
func (c *UserEditController) Post() {
	var input userEditForm
	c.ParseForm(&input)
	c.FormVerify(&input)

	if re := models.UserUpdate(c.GetSession("UID").(int), input.Nickname, input.InShort, input.Sex, input.Birthday, input.Introduce); re == false {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}

// AvatarURLUpdata 修改头像
func (c *UserEditController) AvatarURLUpdata() {
	path := c.GetString("path")

	re := models.AvatarURLUpdate(c.GetSession("UID").(int), path)
	if re == false {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}

// QrImgUpdata 打赏二维码修改
func (c *UserEditController) QrImgUpdata() {
	path := c.GetString("path")

	re := models.QrImgUpdate(c.GetSession("UID").(int), path)
	if re == false {
		c.ResponseJSON(false, "操作失败")
	}

	c.ResponseJSON(true, "操作成功")
}
