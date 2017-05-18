package routers

import (
	"PaperManagementServer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//get参数都在URL后面,post参数都在request body里

	//根据厂名查询实时数据,/factory/:factory?group="一号线"
	beego.Router("/factory/:factory", &controllers.MainController{}, "get:Get")
	//上传实时数据
	beego.Router("/factory/", &controllers.MainController{}, "post:Post")
	//根据厂名查询产线
	beego.Router("/factory/:factory/groups", &controllers.MainController{}, "get:GetGroups")
	//示例：/histroy/:factory?date=2017-05-14&class=D&group=a
	beego.Router("/history/:factory", &controllers.HistoryController{}, "get:Get")
	//上传历史数据
	beego.Router("/history/", &controllers.HistoryController{}, "post:Post")
}
