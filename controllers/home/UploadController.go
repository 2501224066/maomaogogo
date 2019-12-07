package home

import (
	"context"
	"maomaogogo/controllers"
	"maomaogogo/helper"

	"io"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

const (
	// Bucket 篮子
	bucket = "maomaogogo"
	// AccessKey ..
	accessKey = "HOQX1pTD1qYr2ts0Zf18iP3VDZH4ZXWrYpz0FqOK"
	// SecretKey ...
	secretKey = "UnrhYre8Gg6y2oOwD_dI8J3GIvrB_5_oIqEmbzyP"
	// url 外链网址
	url = "http://qiniu.unclear.top"
)

// UploadController 上传控制器
type UploadController struct {
	controllers.Controller
}

// Post ...
func (c *UploadController) Post() {
	f, h, err := c.GetFile("file")
	if err != nil {
		c.ResponseJSON(false, "上传失败")
	}
	defer f.Close()

	file, _ := h.Open() // 这里获得的实际就是一个io,通过源码看到这个open方法最终返回的是一个结构体,其内部包含了 io.Reader的接口

	filename := helper.Md5(h.Filename) // 这里是计算了一个md5(time())的字符串作为文件名

	if filename, err = QiniuUpload(file, h.Size, filename); err != nil { // 通过h.size 即可获得文件大小
		c.ResponseJSON(false, "上传失败")
	} else {
		data := make(map[string]string)
		data["path"] = url + "/" + filename

		c.ResponseJSON(true, "上传成功", data)
	}
}

// QiniuUpload 七牛上传
func QiniuUpload(localFile io.Reader, size int64, filename string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan // 空间对应的机房
	cfg.UseHTTPS = false           // 是否使用https域名
	cfg.UseCdnDomains = false      // 上传是否使用CDN上传加速

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.Put(context.Background(), &ret, upToken, filename, localFile, size, &putExtra)
	if err != nil {
		return "", err
	}

	return ret.Key, nil
}
