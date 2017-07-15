// forceController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type ForceController struct {
	beego.Controller
}

/*
根据厂名获得刷新得数据
*/
func (this *ForceController) Get() {
	factory := this.Ctx.Input.Param(":factory")
	//根据厂名查询强制刷新数据
	var forceData models.ForceData
	forceData.Name = factory
	forceData.Refreshed = true
	forceDatas, err := models.ReadForceData(forceData)
	if err != nil {
		beego.Debug("ForceController", "read foceData failed:"+err.Error())
		this.Ctx.WriteString(err.Error())
		return
	}
	beego.Debug("ForceController", "forceDatas lenght is "+string(len(forceDatas)))
	if len(forceDatas) > 0 {
		this.Data["json"] = &forceDatas[0]
		this.ServeJSON()
	}
}

/*
接收强制刷新数据
*/
func (this *ForceController) Post() {
	//获得requestbody
	requestBody := this.Ctx.Input.RequestBody
	beego.Debug("ForceController", "requestBody:"+string(requestBody))

	var forceData models.ForceData
	err := json.Unmarshal(requestBody, &forceData)
	if err != nil {
		beego.Debug("ForceController", "unmarshal data failed:"+err.Error())
		this.Ctx.WriteString(err.Error())
		return
	}
	forceData.Refreshed = false
	err = models.CreateForceData(forceData)
	if err != nil {
		beego.Debug("ForceController", "create forceData failed:"+err.Error())
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString("{\"response\":\"post succes\"}")
}
