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
	this.Ctx.WriteString("post success")
}
