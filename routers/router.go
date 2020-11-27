package routers

import (
	"url-shortener/controllers"

	beego "github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/:shorturl(^\\w{5}$)", &controllers.RedirectController{})

}
