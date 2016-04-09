package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mozikachu/beego-demo/docs"
	_ "github.com/mozikachu/beego-demo/routers"
)

func init() {
	err := orm.RegisterDataBase("default", "mysql", "root:@/beego_demo_pj?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
