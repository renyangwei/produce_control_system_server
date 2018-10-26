// finishInfoController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"
	"time"

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
	var finishInfo models.FinishInfo
	finishInfo.Cname = finishInfos[0].Cname
	finishInfo.Group = finishInfos[0].Group
	err = models.DeleteFinishInfo(finishInfo)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("FinishInfoController", "delete finishInfo:"+err.Error())
		return
	}
	//插入
	for _, finishInfo = range finishInfos {
		err = models.InsertFinishInfo(finishInfo)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			beego.Error("FinishInfoController", "insert finishInfo:"+err.Error())
			return
		}
	}
	this.Ctx.WriteString("post success")
}

//获得完工资料
func (this *FinishInfoController) Get() {
	var cName = this.GetString("factory", "")
	var group = this.GetString("group", "")
	beego.Debug("cName=", cName)
	if cName == "" {
		this.Ctx.WriteString("parmas is empty")
		beego.Debug("parmas is empty")
		return
	}
	//获得今天的时间
	today := time.Now().Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	var finishInfo models.FinishInfo
	finishInfo.Cname = cName
	finishInfo.Group = group
	finishInfo.StartTime = today
	finishInfo.FinishTime = tomorrow
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
