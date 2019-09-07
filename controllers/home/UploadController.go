package home

import (
	"fmt"
	"maomaogogo/controllers"
)

// UploadController 上传控制器
type UploadController struct {
	controllers.BaseController
}

// Post ...
func (c *UploadController) Post() {
	f, h, err := c.GetFile("file")
	if err != nil {
		fmt.Println(err)
		c.ResponseJSON(false, "上传失败")
	}
	defer f.Close()
	path := "static/upload/" + h.Filename
	c.SaveToFile("file", path)

	data := make(map[string]string)
	data["path"] = "/" + path

	c.ResponseJSON(true, "上传成功", data)
}
