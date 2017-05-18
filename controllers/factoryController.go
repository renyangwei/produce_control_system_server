package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/*
get
通过factory字段查询
*/
func (c *MainController) Get() {
	factory := c.Ctx.Input.Param(":factory")
	group := c.GetString("Group", "一号线")
	beego.Debug("mainController.get, factory:", factory)
	beego.Debug("mainController.get, group:", group)
	if factory == "" {
		c.Ctx.WriteString("factory is empty")
		return
	}

	var fac models.Factory
	fac.Name = factory
	fac.Group = group
	beego.Debug("MainController.Get, before query factory:", fac)
	fac, err := models.FactoryRead(fac)
	if err != nil {
		beego.Error("MainController.Get", err.Error())
		c.Ctx.WriteString(err.Error())
		return
	}
	beego.Debug("MainController.Get, after query factory:", fac)
	c.Data["json"] = &fac
	c.ServeJSON()
}

/*
post
先查询，然后修改
*/
func (c *MainController) Post() {
	//获得requestbody
	requestBody := c.Ctx.Input.RequestBody
	beego.Debug("MainController", "requestBody:"+string(requestBody))

	var fac models.Factory
	err := json.Unmarshal(requestBody, &fac)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		beego.Error("MainController", "unMarshl json:"+err.Error())
		return
	}
	if fac.Group == "" {
		fac.Group = "一号线"
	}
	created, id, err := models.FactoryReadOrCreate(fac)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	beego.Debug("MainController.Post, id:", id)
	beego.Debug("MainController.Post, created:", created)
	if !created {
		fac.Id = id
		err = models.FactoryUpdate(fac)
		if err != nil {
			beego.Error("MainController", "update factory:"+err.Error())
			c.Ctx.WriteString(err.Error())
			return
		}
	}
	c.Ctx.WriteString("post success")
}

/*
获得产线数据
*/
func (this *MainController) GetGroups() {
	var factoryStr = this.Ctx.Input.Param(":factory")
	beego.Debug("MainController.getGroups, factory:", factoryStr)
	var factory models.Factory
	factory.Name = factoryStr
	factories, err := models.ReadFactoryGroups(factory)
	if err != nil {
		beego.Error("MainController.GetGroups, ReadFactoryGroups failed:", err.Error())
		this.Ctx.WriteString(err.Error())
		return
	}
	beego.Debug("MainController.getGroups, factories", factories)
	var groupJson string
	for _, factory = range factories {
		groupJson += factory.Group + ","
	}
	type GroupJson struct {
		Group string
	}
	var gj GroupJson
	gj.Group = groupJson
	beego.Debug("MainController.getGroups, groupJson:" + groupJson)
	this.Data["json"] = &gj
	this.ServeJSON()
}
