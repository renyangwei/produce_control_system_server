// finishInfoController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type FinishInfoController struct {
	beego.Controller
}

//插入完工资料
func (this *FinishInfoController) Post() {
	requestBody := this.Ctx.Input.RequestBody
	beego.Debug("FinishInfoController", "requestBody:"+string(requestBody))
	var finishInfos []models.FinishInfo
	err := json.Unmarshal(requestBody, &finishInfos)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("FinishInfoController", "unMarshl json:"+err.Error())
		return
	}
	//根据cname查询
	var finishInfo models.FinishInfo
	finishInfo.Cname = finishInfos[0].Cname
	finishInfosQuery, err := models.ReadFinishInfo(finishInfo)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("FinishInfoController", "query finishInfo:"+err.Error())
		return
	}
	//删除完工资料
	err = models.DeleteFinishInfo(finishInfosQuery)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("FinishInfoController", "delete finishInfo:"+err.Error())
		return
	}
	//插入
	err = models.InsertFinishInfo(finishInfos)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("FinishInfoController", "insert finishInfo:"+err.Error())
		return
	}
	this.Ctx.WriteString("post success")
}

//获得完工资料
func (this *FinishInfoController) Get() {
	var cName = this.GetString("factory", "")
	var startTime = this.GetString("start_time", "")
	var finishTime = this.GetString("finish_time", "")
	beego.Debug("cName=", cName)
	beego.Debug("startTime=", startTime)
	beego.Debug("finishTime=", finishTime)
	if cName == "" || startTime == "" || finishTime == "" {
		this.Ctx.WriteString("parmas is empty")
		beego.Debug("parmas is empty")
		return
	}
	var finishInfo models.FinishInfo
	finishInfo.Cname = cName
	finishInfo.StartTime = startTime
	finishInfo.FinishTime = finishTime
	//根据厂名和时间查询
	finishInfos, err := models.ReadFinishInfo(finishInfo)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("FinishInfoController", "read finishInfos:"+err.Error())
		return
	}
	beego.Debug("FinishInfoController, read finishInfos:", finishInfos)
	this.Data["json"] = &finishInfos
	this.ServeJSON()
}
