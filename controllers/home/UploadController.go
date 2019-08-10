package home

import (
	"fmt"
	"maomaogogo/controllers"
)

type UploadController struct {
	controllers.BaseController
}

func (this *UploadController) Post() {
	f, h, err := this.GetFile("file")
	if err != nil {
		fmt.Println(err)
		this.ResponseJson(false, "上传失败")
	}
	defer f.Close()
	path := "static/upload/" + h.Filename
	this.SaveToFile("file", path)

	data := make(map[string]string)
	data["path"] = "/" + path

	this.ResponseJson(true, "上传成功", data)
}
