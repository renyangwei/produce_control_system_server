// searchController
package controllers

import (
	"PaperManagementServer/models"
	"encoding/json"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type SearchController struct {
	beego.Controller
}

/*
订单数据
*/
type OrderData struct {
	Scxh        string    `json:"Scxh"`   //序号
	Mxbh        string    `json:"Mxbh"`   //订单号
	Khjc        string    `json:"Khjc"`   //客户简称
	Zbdh        string    `json:"Zbdh"`   //材质
	Klzhdh      string    `json:"Klzhdh"` //楞别
	Zd          string    `json:"Zd"`     //纸度
	Zbcd        string    `json:"Zbcd"`   //切长
	Pscl        string    `json:"Pscl"`   //排产数量
	Ddms        string    `json:"Ddms"`   //留言
	Zt          string    `json:"Zt"`     //是否正在进行
	Ks          string    `json:"Ks"`     //剖
	Sm2         string    `json:"Sm2"`
	Zbcd2       string    `json:"Zbcd2"` //切长
	Xbmm        string    `json:"Xbmm"`
	Scbh        string    `json:"Scbh"`
	Ms          string    `json:"Ms"`
	Finish_time time.Time `json:"FinishTime"` //预计完工时间
}

/*
完工数据
*/
type FinishInfoData struct {
	Mxbh  string `json:"Mxbh"`  //订单号
	Khjc  string `json:"Khjc"`  //客户简称
	Zbdh  string `json:"Zbdh"`  //材质
	Zbkd  string `json:"Zbkd"`  //纸板宽
	Hgpsl string `json:"Hgpsl"` //合格数
	Blpsl string `json:"Blpsl"` //不良数
	Pcsl  string `json:"Pcsl"`  //排产数
	Zbcd  string `json:"Zbcd"`  //切长
}

/*
查询接口
*/
func (this *SearchController) Get() {
	var searchType = this.GetString("type", "")
	beego.Debug("searchController, type:", searchType)
	var data = this.GetString("data", "")
	beego.Debug("SearchController, data:", data)
	var cname = this.GetString("cname", "")
	beego.Debug("SearchController, cname:", cname)
	var startTime = this.GetString("start_time", "")
	beego.Debug("SearchController, startTime:", startTime)
	var finishTime = this.GetString("finish_time", "")
	beego.Debug("SearchController, finishTime:", finishTime)
	var err error

	switch searchType {
	case "order":
		//返回搜索结果
	case "finish_info":
		var finishInfos []models.FinishInfo
		if startTime == "" {
			finishInfos, err = models.SearchFinishInfo(cname)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				beego.Error("searchFinishInfo, searchFinishInfo err:", err.Error())
				return
			}
		} else {
			var finishInfo models.FinishInfo
			finishInfo.Cname = cname
			finishInfo.StartTime = startTime
			finishInfo.FinishTime = finishTime
			finishInfos, err = models.SearchFinishInfoByTime(finishInfo)
			if err != nil {
				this.Ctx.WriteString(err.Error())
				beego.Error("searchFinishInfo, SearchFinishInfoByTime err:", err.Error())
				return
			}
		}
		beego.Debug("SearchController, searchFinishInfos:", finishInfos)
		var respFinishInfo []models.FinishInfo
		for _, finishInfo := range finishInfos {
			var finishInfoData FinishInfoData
			err := json.Unmarshal([]byte(finishInfo.Data), &finishInfoData)
			if err != nil {
				beego.Error("SearchController, unmarshal finishInfo err:", err.Error())
				continue
			}
			if strings.Contains(finishInfoData.Mxbh, data) || strings.Contains(finishInfoData.Khjc, data) || strings.Contains(finishInfoData.Zbcd, data) {
				respFinishInfo = append(respFinishInfo, finishInfo)
			}
		}
		beego.Debug("SearchController, respFinishInfo:", respFinishInfo)
		this.Data["json"] = &respFinishInfo
		this.ServeJSON()
	}
}

/*
上传需要查询的参数
*/
func (this *SearchController) Post() {
	requestBody := this.Ctx.Input.RequestBody
	beego.Debug("SearchController", "requestBody:"+string(requestBody))
	var searchRequest models.SearchRequest
	err := json.Unmarshal(requestBody, &searchRequest)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("SearchController, unmarshal requestbody err:", err.Error())
		return
	}
	searchRequest.IsSearched = false
	err = models.InsertSearchRequest(searchRequest)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		beego.Error("SearchController, insert SearchRequest err:", err.Error())
		return
	}
	this.Ctx.WriteString("post success")
}
