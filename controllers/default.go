package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"onework/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	//1.获取请求数据
	name:=c.Ctx.Input.Query("name")
	age:=c.Ctx.Input.Query("age")

	//2.使用固定数据进行数据校验
	if name!="l" && age!="19"{
		//代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误"))
		return
	}else {
		c.Ctx.ResponseWriter.Write([]byte("恭喜，数据校验正确"))

	}




	c.Data["Website"] = "Tom-ly.me"
	c.Data["Email"] = "3419572132@gmail.com"
	c.TplName = "index.tpl"
}
//func (c*MainController)Post()  {
	//1.解析post请求的参数
//	name :=c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	fmt.Println(name,age)

	//2.进行数据校验
//	if name!="adam"&&age!="19"{
//		c.Ctx.ResponseWriter.Write([]byte("数据校验失败"))
//		return
//	}else {
//		c.Ctx.WriteString("数据校验成功")
//	}


//}
func (c*MainController)Post(){
	//1.解析post请求的参数jsong格式


	var person models.Person
	dataBytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err!=nil{
		c.Ctx.WriteString("数据接收错误")
		return
	}
	err=json.Unmarshal(dataBytes,&person)
	if err!=nil{
		c.Ctx.WriteString("数据解析失败")
		return
	}

	fmt.Println("姓名:",person.Name)
	fmt.Println("年龄:",person.Birthday)
	fmt.Println("性别:",person.Addrss)
	fmt.Println("性别:",person.Nick)


	c.Ctx.WriteString("数据解析成功")

}