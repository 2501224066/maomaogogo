package routers

import (
	"maomaogogo/controllers/errors"
	"maomaogogo/controllers/home"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &home.IndexController{})
	beego.Router("/404", &errors.ErrorController{}, "get:Error404")
	beego.Router("/send_email/?:type", &home.SendEmailController{})
	beego.Router("/upload", &home.UploadController{})

	beego.Router("/register", &home.RegisterController{})
	beego.Router("/login", &home.LoginController{})
	beego.Router("/logout", &home.LogoutController{})

	beego.Router("/user_read", &home.UserReadController{})
	beego.Router("/user_edit", &home.UserEditController{})
	beego.Router("/avatar_url_update", &home.UserEditController{}, "post:AvatarUrlUpdata")
	beego.Router("/qr_img_update", &home.UserEditController{}, "post:QrImgUpdata")

	beego.Router("/article_create", &home.ArticleCreateController{})
	beego.Router("/article_edit/?:article_id", &home.ArticleEditController{})
	beego.Router("/article_list/?:node_id", &home.ArticleListController{})
	beego.Router("/article_read/?:article_id", &home.ArticleReadController{})

	beego.Router("/user/article", &home.UserResourceController{}, "get:Article")

	beego.Router("/user_op/article_like/?:article_id", &home.UserOpController{}, "get:ArticleLike")
	beego.Router("/user_op/article_collect/?:article_id", &home.UserOpController{}, "get:ArticleCollect")
	beego.Router("/user_op/article_del/?:article_id", &home.UserOpController{}, "get:ArticleDel")
}
