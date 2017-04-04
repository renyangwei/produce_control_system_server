package main

import (
	_ "PaperManagementServer/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/managementServer.log"}`)
	beego.Run()
}
