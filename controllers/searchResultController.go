// searchResultController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

type SearchResultController struct {
	beego.Controller
}

/*
获取搜索结果,App端调用
*/
func (this *SearchResultController) Get() {
	cname := this.GetString("cname", "")
	group := this.GetString("group", "")
	typee := this.GetString("type", "")
	if cname == "" || group == "" || typee == "" {
		beego.Error("param is empty")
		this.Ctx.WriteString("param is empty")
		return
	}
	//根据参数获取搜索结果
	var searchResult models.SearchResult
	searchResult.Cname = cname
	searchResult.Group = group
	searchResult.Type = typee
	searchResults, err := models.ReadSearchResult(searchResult)
	if err != nil {
		beego.Error("read search result err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
	beego.Debug("SearchResultController, search results:", searchResults)
	this.Data["json"] = &searchResults
	this.ServeJSON()
	//删除
	err = models.DeleteSearchResult(searchResult)
	if err != nil {
		beego.Error("delete search result err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
}

/*
上传搜索结果
*/
func (this *SearchResultController) Post() {
	requestBody := this.Ctx.Input.RequestBody
	beego.Debug("SearchResultController", "requestBody:"+string(requestBody))
	//解析
	var searchResults []models.SearchResult
	err := json.Unmarshal(requestBody, &searchResults)
	if err != nil {
		beego.Error("SearchResultController, unmarshl params err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
	for i, _ := range searchResults {
		beego.Debug("SearchResultController, startTime:", searchResults[i].StartTime)
		if searchResults[i].StartTime == "" {
			searchResults[i].Type = "order"
		} else {
			searchResults[i].Type = "finish_info"
		}
	}
	//保存搜索结果
	err = models.InsertSearchResult(searchResults)
	if err != nil {
		beego.Error("SearchResultController, insert search result err:", err)
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString("post success")
}
