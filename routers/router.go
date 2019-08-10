package routers

import (
	"maomaogogo/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &home.IndexController{})
	beego.Router("/upload", &home.UploadController{})
	beego.Router("/register", &home.RegisterController{})
	beego.Router("/login", &home.LoginController{})
	beego.Router("/logout", &home.LogoutController{})
	beego.Router("/info", &home.InfoController{})
	beego.Router("/info_update", &home.InfoUpdateController{})
	beego.Router("/avatar_url_update", &home.AvatarUrlUpdateController{})
	beego.Router("/send_email/?:type", &home.SendEmailController{})
	beego.Router("/create_article", &home.CreateArticleController{})
	beego.Router("/read_article/?:id", &home.ReadArticleController{})
}
