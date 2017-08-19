// paper
package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//工厂
type Factory struct {
	Id    int64  `orm:"pk;auto"`
	Name  string `json:"Factory"` //厂名
	Other string `json:"Other"`   //实时数据
	Group string `json:"Group"`   //产线
}

//历史数据
type History struct {
	Id    int64  `orm:"pk;auto"`
	Name  string `json:"Factory"` //厂名
	Other string `json:"Other"`   //历史数据
	Class string `json:"Class"`   //班组
	Date  string `json:"Time"`    //日期
	Group string `json:"Group"`   //产线
}

//强制刷新
type ForceData struct {
	Id        int64  `orm:"pk;auto"`
	Name      string `json:"Factory"`   //厂名
	Class     string `json:"Class"`     //班组
	Date      string `json:"Time"`      //日期
	Group     string `json:"Group"`     //产线
	Refreshed bool   `json:"Refreshed"` //是否刷新过
}

//完工资料
type FinishInfo struct {
	Id         int64  `orm:"pk;auto"`     //Id
	Cname      string `json:"Cname"`      //公司名
	Data       string `json:"Data"`       //数据
	StartTime  string `json:"StartTime"`  //开始时间
	FinishTime string `json:"FinishTime"` //完成时间
	Group      string `json:"Group"`      //产线
}

//订单信息
type Order struct {
	Id    int64  `orm:"pk;auto"` //Id
	Cname string `json:"Cname"`  //公司名
	Data  string `json:"Data"`   //数据
	Group string `json:"Group"`  //产线
}

/*
搜索请求参数
*/
type SearchRequest struct {
	Id         int64  `orm:"pk;auto"`     //Id
	Cname      string `json:"Cname"`      //公司名
	Data       string `json:"Data"`       //搜索内容
	StartTime  string `json:"StartTime"`  //开始时间
	FinishTime string `json:"FinishTime"` //完成时间
	Group      string `json:"Group"`      //产线
	Type       string `json:"Type"`       //类型，order、finish_info
	IsSearched bool   `json:"IsSearched"` //是否搜索过
}

/*
搜索结果
*/
type SearchResult struct {
	Id         int64  `orm:"pk;auto"`     //Id
	Cname      string `json:"Cname"`      //公司名
	Data       string `json:"Data"`       //数据
	StartTime  string `json:"StartTime"`  //开始时间
	FinishTime string `json:"FinishTime"` //完成时间
	Group      string `json:"Group"`      //产线
	Type       string `json:"Type"`       //类型，order、finish_info
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/paper_management?charset=utf8", 30)
	orm.RegisterModel(new(Factory))
	orm.RegisterModel(new(History))
	orm.RegisterModel(new(ForceData))
	orm.RegisterModel(new(FinishInfo))
	orm.RegisterModel(new(Order))
	orm.RegisterModel(new(SearchRequest))
	orm.RegisterModel(new(SearchResult))
	orm.RunSyncdb("default", false, true)
}

/*
根据厂名和产线查询
*/
func FactoryRead(factory Factory) (fac Factory, err error) {
	o := orm.NewOrm()
	err = o.Read(&factory, "Name", "Group")
	return factory, err
}

/*
根据厂名和产线查询
如果没有就创建
*/
func FactoryReadOrCreate(factory Factory) (created bool, id int64, err error) {
	o := orm.NewOrm()
	created, id, err = o.ReadOrCreate(&factory, "Name", "Group")
	return created, id, err
}

/*
修改数据
*/
func FactoryUpdate(factory Factory) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&factory)
	return err
}

/*
根据厂名查询产线
*/
func ReadFactoryGroups(factory Factory) (factories []Factory, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("Factory").Filter("Name", factory.Name).All(&factories)
	return
}

/*
根据厂名、班组、日期和产线查询
*/
func ReadHistory(history History) (his History, err error) {
	beego.Debug("ReadHistory:", history)
	o := orm.NewOrm()
	err = o.Read(&history, "Name", "Class", "Date", "Group")
	return history, err
}

/*
保存数据
*/
func CreateHistory(history History) (created bool, id int64, err error) {
	o := orm.NewOrm()
	created, id, err = o.ReadOrCreate(&history, "Name", "Class", "Date")
	return created, id, err
}

/*
修改历史数据
*/
func UpdateHistory(history History) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&history)
	return
}

/*
保存强制刷新的数据
*/
func CreateForceData(forceData ForceData) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&forceData)
	return
}

/*
获得强制刷新数据
*/
func ReadForceData(forceData ForceData) (forceD []ForceData, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("ForceData").Filter("Name", forceData.Name).Filter("Refreshed", forceData.Refreshed).All(&forceD)
	return forceD, err
}

/*
更新是否刷新过
*/
func UpdateForceData(forceData ForceData) (err error) {
	beego.Debug("UpdateForceData", forceData)
	o := orm.NewOrm()
	_, err = o.Update(&forceData, "Refreshed")
	return
}

/*
获得历史数据的产线
*/
func ReadHistoryGorup(history History) (his []History, err error) {
	beego.Debug("ReadHistoryGorup", history)
	o := orm.NewOrm()
	_, err = o.QueryTable("History").Filter("Name", history.Name).GroupBy("Group").OrderBy("Group").All(&his, "Group")
	//	_, err = o.Raw("select group from history where name = ? order by decode(group, '一号线', 1, '二号线', 2, '三号线', 3)", history.Name).QueryRows(&his)
	return his, err
}

/*
获得历史数据的班组
*/
func ReadHistoryClass(history History) (his []History, err error) {
	beego.Debug("ReadFactoryClass", history)
	o := orm.NewOrm()
	_, err = o.QueryTable("History").Filter("Name", history.Name).GroupBy("Class").All(&his, "Class")
	return his, err
}

/*
获得最近一次历史数据
*/
func ReadLastHistory(history History) (his []History, err error) {
	beego.Debug("ReadLastHistory", history)
	o := orm.NewOrm()
	_, err = o.QueryTable("History").Filter("Name", history.Name).OrderBy("-Date").All(&his, "Id", "Name", "Date", "Class", "Group")
	return his, err
}

//插入订单
func InsertOrder(order Order) (err error) {
	beego.Debug("InsertOrder", order)
	o := orm.NewOrm()
	_, err = o.Insert(&order)
	return
}

//读取订单
func ReadOrder(order Order) (orders []Order, err error) {
	beego.Debug("ReadOrder", order)
	o := orm.NewOrm()
	_, err = o.QueryTable("Order").Filter("Cname", order.Cname).Filter("Group", order.Group).All(&orders)
	return
}

//删除订单
func DeleteOrder(order Order) (err error) {
	beego.Debug("DeleteOrder", order)
	o := orm.NewOrm()
	_, err = o.QueryTable("Order").Filter("Cname", order.Cname).Delete()
	return
}

//插入完工资料
func InsertFinishInfo(finishInfo FinishInfo) (err error) {
	beego.Debug("InsertFinishInfo", finishInfo)
	o := orm.NewOrm()
	_, err = o.Insert(&finishInfo)
	return
}

//读取完工资料
func ReadFinishInfo(finishInfo FinishInfo) (finishInfos []FinishInfo, err error) {
	beego.Debug("ReadFinishInfo", finishInfo)
	o := orm.NewOrm()
	_, err = o.QueryTable("FinishInfo").Filter("Cname", finishInfo.Cname).Filter("StartTime__gte", finishInfo.StartTime).Filter("FinishTime__lte", finishInfo.FinishTime).Filter("Group", finishInfo.Group).All(&finishInfos)
	return
}

//删除完工资料
func DeleteFinishInfo(finishInfo FinishInfo) (err error) {
	beego.Debug("DeleteFinishInfo", finishInfo)
	o := orm.NewOrm()
	_, err = o.QueryTable("FinishInfo").Filter("Cname", finishInfo.Cname).Delete()
	return
}

/*
搜索订单
*/
func SearchOrder(cname string) (orders []Order, err error) {
	beego.Debug("SearchOrder, cname:", cname)
	o := orm.NewOrm()
	_, err = o.QueryTable("Order").Filter("Cname", cname).All(&orders)
	return
}

/*
根据时间段搜索完工资料
*/
func SearchFinishInfoByTime(finishInfo FinishInfo) (finishInfos []FinishInfo, err error) {
	beego.Debug("SearchFinishInfoByTime", finishInfo)
	o := orm.NewOrm()
	_, err = o.QueryTable("FinishInfo").Filter("Cname", finishInfo.Cname).Filter("StartTime__gte", finishInfo.StartTime).Filter("FinishTime__lte", finishInfo.FinishTime).All(&finishInfos)
	return
}

/*
搜索所有时间段的完工资料
*/
func SearchFinishInfo(cname string) (finishInfos []FinishInfo, err error) {
	beego.Debug("SearchFinishInfo", cname)
	o := orm.NewOrm()
	_, err = o.QueryTable("FinishInfo").Filter("Cname", cname).All(&finishInfos)
	return
}

/*
保存搜索请求
*/
func InsertSearchRequest(searchRequest SearchRequest) (err error) {
	beego.Debug("InsertSearchRequest", searchRequest)
	o := orm.NewOrm()
	_, err = o.Insert(&searchRequest)
	return
}

/*
读取搜索请求
*/
func ReadSearchRequest(searchRequest SearchRequest) (searchRequests []SearchRequest, err error) {
	beego.Debug("ReadSearchRequest", searchRequest)
	o := orm.NewOrm()
	_, err = o.QueryTable("SearchRequest").Filter("Cname", searchRequest.Cname).Filter("Group", searchRequest.Group).Filter("IsSearched", searchRequest.IsSearched).All(&searchRequests)
	return

}
