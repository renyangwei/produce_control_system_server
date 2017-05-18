package main

import (
	_ "PaperManagementServer/routers"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/orm"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/managementServer.log"}`)
	beego.Run()
}
