package routers

import (
	"PaperManagementServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/factory/:factory", &controllers.MainController{}, "get:Get")
	beego.Router("/factory/", &controllers.MainController{}, "post:Post")
}
