// paper
package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

//工厂
type Factory struct {
	Id    int64  `orm:"pk;auto"`
	Name  string `json:"Factory"` //厂名
	Other string `json:"Other"`
	Group string `json:"Group"` //产线
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:123456@/paper_management?charset=utf8", 30)
	orm.RegisterModel(new(Factory))
	orm.RunSyncdb("default", false, true)
}

/*
根据厂名查询
*/
func FactoryRead(factory Factory) (fac Factory, err error) {
	o := orm.NewOrm()
	err = o.Read(&factory, "Name")
	return factory, err
}

/*
根据厂名查询
如果没有就创建
*/
func FactoryReadOrCreate(factory Factory) (created bool, id int64, err error) {
	o := orm.NewOrm()
	created, id, err = o.ReadOrCreate(&factory, "Name")
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
