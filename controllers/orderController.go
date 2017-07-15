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
	//根据cname查询
	var order models.Order
	order.Cname = orders[0].Cname
	ordersQuery, err := models.ReadOrder(order)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("OrderController", "query order by cname:"+err.Error())
		return
	}
	beego.Debug("OrderController, query order by cname:", ordersQuery)
	//删除
	err = models.DeleteOrder(ordersQuery)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("OrderController", "delete orders:"+err.Error())
		return
	}
	//插入
	err = models.InsertOrder(orders)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("OrderController", "insert orders:"+err.Error())
		return
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
	var order models.Order
	order.Cname = cName
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
