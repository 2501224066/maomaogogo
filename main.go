package main

import (
	"maomaogogo/conf"
	"maomaogogo/controllers/errors"
	"maomaogogo/helper"
	_ "maomaogogo/routers"

	"github.com/astaxie/beego"
)

func init() {
	conf.InitSession()
	beego.ErrorController(&errors.ErrorController{})
	beego.AddFuncMap("isActive", helper.IsActive)
	beego.AddFuncMap("StrSplitForBlankSpace", helper.StrSplitForBlankSpace)
	beego.AddFuncMap("TimeInterval", helper.TimeInterval)
}

func main() {
	beego.Run()
}
