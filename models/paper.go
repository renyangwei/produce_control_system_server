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

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/paper_management?charset=utf8", 30)
	orm.RegisterModel(new(Factory))
	orm.RegisterModel(new(History))
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
