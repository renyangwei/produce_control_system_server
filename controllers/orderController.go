// orderController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

//插入订单
func (this *OrderController) Post() {
	requestBody := this.Ctx.Input.RequestBody
	beego.Debug("OrderController", "requestBody:"+string(requestBody))
	var orders []models.Order
	err := json.Unmarshal(requestBody, &orders)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("OrderController", "unMarshl json:"+err.Error())
		return
	}
	//删除
	var order models.Order
	order.Cname = orders[0].Cname
	err = models.DeleteOrder(order)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("OrderController", "delete order:"+err.Error())
		return
	}
	//插入
	for _, order = range orders {
		err = models.InsertOrder(order)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			beego.Error("OrderController", "insert order:"+err.Error())
			return
		}
	}
	this.Ctx.WriteString("post success")
}

//获取订单
func (this *OrderController) Get() {
	var cName = this.GetString("factory", "")
	if cName == "" {
		this.Ctx.WriteString("parmas is empty")
		beego.Debug("parmas is empty")
		return
	}
	var group = this.GetString("group", "")
	var order models.Order
	order.Cname = cName
	order.Group = group
	//根据厂名获取订单
	orders, err := models.ReadOrder(order)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("OrderController", "read orders:"+err.Error())
		return
	}
	beego.Debug("OrderController, read orders:", orders)
	this.Data["json"] = &orders
	this.ServeJSON()
}
