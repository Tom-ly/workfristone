package db_mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"onework/models"
)


var Db*sql.DB
//在初始化函数中连接数据库

func init()  {
	fmt.Println("连接mysql数据库")
	config :=beego.AppConfig
	dbDriver :=config.String("ddb_driverName")
	dbUser :=config.String("db_User")
	dbPassword :=config.String("db_password")
	dbIp :=config.String("db_ip")
	dbName :=config.String("db_name")

	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db,err:=sql.Open(dbDriver,connUrl)
	if err!=nil{
		panic("数据库连接错误，请检查")
	}
	//为全局变量赋值
	Db=db



}


//将用户信息保存到数据库中
func InsertUser(user models.User) (int64,error) {


	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名:",user.Nick,"密码:",user.Password)
	result,err := Db.Exec("insert into user(nick, password) values(?,?)",user.Nick,user.Password)
	if err != nil {
		//保存数据时遇到错误
		return -1,err
	}

	id, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return id,err

}


//查询用户
func QueryUser()  {
	Db.QueryRow("select *from ")


}