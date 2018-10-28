// searchRequestController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type SearchRequestController struct {
	beego.Controller
}

type SearchRequestStruct struct {
	Type  string //搜索类型:order,finish_info
	Cname string //公司名
	Group string //产线
	Info  string //搜索内容
}

/*
获取搜索参数
*/
func (this *SearchRequestController) Get() {
	cname := this.GetString("cname", "")
	beego.Debug("SearchRequestController.Get, cname:", cname)
	if cname == "" {
		beego.Error("cname is empty")
		this.Ctx.WriteString("cname is empty")
		return
	}
	//根据公司名和是否搜索过查询
	var searchRequest models.SearchRequest
	searchRequest.Cname = cname
	searchRequest.IsSearched = false
	searchRequests, err := models.ReadSearchRequest(searchRequest)
	if err != nil {
		beego.Error("SearchRequestController, read search request err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
	beego.Debug("SearchRequestController, search requests:", searchRequests)
	this.Data["json"] = &searchRequests
	this.ServeJSON()
	//修改是否搜索过
	searchRequests.IsSearched = true
	err = models.UpdateSearchRequest(searchRequests)
	if err != nil {
		beego.Error("SearchRequestController, update search request err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
}

/*
上传搜索参数, APP端调用
*/
func (this *SearchRequestController) Post() {
	requestBody := this.Ctx.Input.RequestBody
	beego.Debug("SearchRequestController", "requestBody:"+string(requestBody))
	var searchRequest models.SearchRequest
	err := json.Unmarshal(requestBody, &searchRequest)
	if err != nil {
		beego.Error("SearchRequestController, unmarshal requestbody err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
	if searchRequest.Cname == "" || searchRequest.Group == "" || searchRequest.Type == "" {
		beego.Debug("SearchRequestController, params is empty")
		this.Ctx.WriteString("params is empty")
		return
	}
	//保存参数
	err = models.InsertSearchRequest(searchRequest)
	if err != nil {
		beego.Error("SearchRequestController, insert search request err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString("{\"response\":\"post succes\"}")
}
