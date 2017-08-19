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
	//查询历史数据,示例：/histroy/:factory?date=2017-05-14&class=D&group=a
	beego.Router("/history/:factory", &controllers.HistoryController{}, "get:Get")
	//根据厂名查询历史数据的产线
	beego.Router("/history/:factory/groups", &controllers.HistoryController{}, "get:GetGroups")
	//根据厂名查询历史数据的班组
	beego.Router("/history/:factory/class", &controllers.HistoryController{}, "get:GetClass")
	//根据厂名查询最近的一次历史数据的时间，班组和生产线
	beego.Router("/history/:factory/last", &controllers.HistoryController{}, "get:GetLast")
	//上传历史数据
	beego.Router("/history/", &controllers.HistoryController{}, "post:Post")
	//获得强制刷新数据，厂名、班组、文件名和日期
	beego.Router("/force/:factory", &controllers.ForceController{}, "get:Get")
	beego.Router("/force/", &controllers.ForceController{}, "post:Post")
	//订单信息
	beego.Router("/order", &controllers.OrderController{}, "get:Get;post:Post")
	//完工信息
	beego.Router("/finish_info", &controllers.FinishInfoController{}, "get:Get;post:Post")
	//搜索接口,
	// /search?type=order&cname=xxx&data=abc&group=一号线
	beego.Router("/search", &controllers.SearchController{}, "get:Get;post:Post")

}
