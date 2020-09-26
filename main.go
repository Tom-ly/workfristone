package main

import (
	"github.com/astaxie/beego"
	_ "onework/db_mysql"
	_ "onework/routers"
)

func main() {
	//config :=beego.AppConfig
	//appName:=config.String("appname")
	//fmt.Println("应用名称：",appName)

	//port,err:=config.Int("httpport")
	//if err!=nil{
	//	panic("项目配置文件解析失败，请检查配置文件")
	//}
	//fmt.Println("端口：",port)


	beego.Run()
}

