package controllers

import (
	"html/template"
	"maomaogogo/models"
	"reflect"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// Controller 基础控制器
type Controller struct {
	beego.Controller
}

// Prepare 初始
func (c *Controller) Prepare() {
	// 路径
	c.Data["Path"] = c.Ctx.Request.RequestURI

	// CSRF
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	// 登录信息
	var (
		LoginStatus       bool
		LoginUser         models.User
		UnreadNoticeCount int64
	)
	uid := c.GetSession("UID")
	if uid != nil {
		LoginStatus = true
		LoginUser = models.UserRead(uid.(int))
		UnreadNoticeCount = models.UnreadNoticeCount(LoginUser.UserID, 1)
	} else {
		LoginStatus = false
	}

	c.Data["LoginStatus"] = LoginStatus
	c.Data["LoginUser"] = LoginUser
	c.Data["PageNum"], _ = beego.AppConfig.Int("page::num")
	c.Data["UnreadNoticeCount"] = UnreadNoticeCount

	c.Layout = "home/layout/layout.html"
}

// ResponseJSON 返回JSON
func (c *Controller) ResponseJSON(isSuccess bool, msg string, data ...interface{}) {
	status := 0
	if isSuccess {
		status = 1
	}
	ret := map[string]interface{}{"status": status, "msg": msg}
	if len(data) > 0 {
		ret["data"] = data[0]
	}
	c.Data["json"] = ret
	c.ServeJSON()
	c.StopRun()
}

// FormVerify 表单校验
func (c *Controller) FormVerify(input interface{}) {
	setVerifyMessage()

	valid := validation.Validation{}
	b, _ := valid.Valid(input)
	if !b {
		arr := strings.Split(valid.Errors[0].Key, ".") //切分例如Password.MinSize
		st := reflect.TypeOf(input).Elem()             // 反射获取 input 信息
		field, _ := st.FieldByName(arr[0])             // 获取 Password 参数信息

		c.ResponseJSON(false, field.Tag.Get("chn")+valid.Errors[0].Message)
	}
}

// setVerifyMessage 设置表单验证messages
func setVerifyMessage() {
	var MessageTmpls = map[string]string{
		"Required":     "不能为空",
		"Min":          "最小值 为 %d",
		"Max":          "最大值 为 %d",
		"Range":        "范围 为 %d 到 %d",
		"MinSize":      "最短长度 为 %d",
		"MaxSize":      "最大长度 为 %d",
		"Length":       "长度必须 为 %d",
		"Alpha":        "必须是有效的字母",
		"Numeric":      "必须是有效的数字",
		"AlphaNumeric": "必须是有效的字母或数字",
		"Match":        "必须匹配 %s",
		"NoMatch":      "必须不匹配 %s",
		"AlphaDash":    "必须是有效的字母、数字或连接符号(-_)",
		"Email":        "必须是有效的电子邮件地址",
		"IP":           "必须是有效的IP地址",
		"Base64":       "必须是有效的base64字符",
		"Mobile":       "必须是有效的手机号码",
		"Tel":          "必须是有效的电话号码",
		"Phone":        "必须是有效的电话或移动电话号码",
		"ZipCode":      "必须是有效的邮政编码",
	}

	validation.SetDefaultMessage(MessageTmpls)
}
