package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"onework/db_mysql"
	"onework/models"

)

type RegisterController struct {
	beego.Controller
}


func (r*RegisterController)Post(){
	//1.解析post请求的参数jsong格式




	var user models.User
	body,err:=ioutil.ReadAll(r.Ctx.Request.Body)//读请求数据
	if err!=nil{
		r.Ctx.WriteString("数据接收错误")
		return
	}
	err=json.Unmarshal(body,&user)//解析数据，并判断是否解析1成功
	if err!=nil{
		r.Ctx.WriteString("数据解析失败")
		return
	}
	r.Ctx.WriteString("数据解析成功")


	fmt.Println("姓名:",user.User)
	fmt.Println("生日:",user.Birthday)
	fmt.Println("住址:",user.Address)
	fmt.Println("称号:",user.Nick)





	//将解析到的用户数据，保存到数据库
	row,err :=db_mysql.InsertUser(user)
	if err!=nil{
		r.Ctx.WriteString("用户注册失败")
		return
	}
	fmt.Println(row)

	result:=models.ResponseResult{
		Code: 0,
		Message: "用户注册成功",
		Date: nil,
	}
	r.Data["json"]=&result
	r.ServeJSON()

}
