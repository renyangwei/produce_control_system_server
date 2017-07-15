// histroyController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type HistoryController struct {
	beego.Controller
}

/*
读取历史数据
*/
func (this *HistoryController) Get() {
	var factory = this.Ctx.Input.Param(":factory")
	var dateStr = this.GetString("date", "")
	var class = this.GetString("class", "")
	var group = this.GetString("group", "")
	beego.Debug("HistoryController.Get:", "factory=", factory, "date=", dateStr, "class=", class, "group=", group)
	if factory == "" || dateStr == "" || class == "" || group == "" {
		this.Ctx.WriteString("parmas is empty")
		beego.Debug("parmas is empty")
		return
	}
	var history models.History
	history.Name = factory
	history.Class = class
	history.Date = dateStr
	history.Group = group
	history, err := models.ReadHistory(history)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("HistoryController.Get:", "read history failed")
		return
	}
	beego.Debug("HistoryController.Get:", history)
	this.Data["json"] = &history
	this.ServeJSON()
}

/*
保存历史数据
*/
func (this *HistoryController) Post() {
	//获得requestbody
	requestBody := this.Ctx.Input.RequestBody
	beego.Debug("HistoryController", "requestBody:"+string(requestBody))
	var history models.History
	err := json.Unmarshal(requestBody, &history)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("HistoryController", "unMarshl json:"+err.Error())
		return
	}
	created, id, err := models.CreateHistory(history)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("HistoryController.Post,", "create history err")
		return
	}
	beego.Debug("HistoryController.Post, id = ", id)
	beego.Debug("HistoryController.Post, created = ", created)
	if !created {
		history.Id = id
		err = models.UpdateHistory(history)
		if err != nil {
			beego.Error("HistoryController.Post", "update factory:"+err.Error())
			this.Ctx.WriteString(err.Error())
			return
		}
	}
	//判断是否强制刷新
	//根据厂名查询
	var forceData models.ForceData
	forceData.Name = history.Name
	forceData.Refreshed = false
	forceDatas, err := models.ReadForceData(forceData)
	beego.Debug("HistoryController.Post, forceDatas is", forceDatas)
	if len(forceDatas) > 0 {
		this.Data["json"] = &forceDatas[0]
		this.ServeJSON()
		//还要修改是否刷新过
		forceDatas[0].Refreshed = true
		err = models.UpdateForceData(forceDatas[0])
		if err != nil {
			beego.Error("HistoryController.Post", "update forceData:"+err.Error())
			this.Ctx.WriteString(err.Error())
			return
		}
		return
	}
	this.Ctx.WriteString("post success")
}

/*
获得历史数据中的产线
*/
func (this *HistoryController) GetGroups() {
	name := this.Ctx.Input.Param(":factory")
	beego.Debug("HistoryController.GetGroups", name)
	var history models.History
	history.Name = name
	//查询产线
	his, err := models.ReadHistoryGorup(history)
	if err != nil {
		beego.Error("HistoryController.GetGroups, readHistoryGorup failed", err.Error())
		return
	}
	beego.Debug("HistoryController.GetGroups, his:", his)
	this.Data["json"] = &his
	this.ServeJSON()
}

/*
获得历史数据中的班组
*/
func (this *HistoryController) GetClass() {
	name := this.Ctx.Input.Param(":factory")
	beego.Debug("HistoryController.GetClass", name)
	var history models.History
	history.Name = name
	//查询产线
	his, err := models.ReadHistoryClass(history)
	if err != nil {
		beego.Error("HistoryController.GetClass, readHistoryGorup failed", err.Error())
		return
	}
	beego.Debug("HistoryController.GetGroups, his:", his)
	this.Data["json"] = &his
	this.ServeJSON()
}

/*
根据厂名查询最近的一次历史数据的日期、班组和生产线
*/
func (this *HistoryController) GetLast() {
	name := this.Ctx.Input.Param(":factory")
	beego.Debug("HistoryController.GetLast", name)
	var history models.History
	history.Name = name
	his, err := models.ReadLastHistory(history)
	if err != nil {
		beego.Error("HistoryController.GetLast, ReadLastHistory failed", err.Error())
		return
	}
	beego.Debug("HistoryController.GetLast, his:", his)
	this.Data["json"] = &his[0]
	this.ServeJSON()
}
