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
	dbDriver :=config.String("db_driverName")
	dbUser :=config.String("db_user")
	dbPassword :=config.String("db_password")
	dbIp :=config.String("db_ip")
	dbName :=config.String("db_name")
	//连接数据库，用sql.Open（）函数，该函数有两个参数
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

	//先将密码进行哈希计算，数据库保存哈希值
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Password))//字符串转成字节切片再进行md5计算
	bytes := hashMd5.Sum(nil)
	user.Password = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名:",user.Nick,"密码:",user.Password)

	//保存到数据库中的函数
	result,err := Db.Exec("insert into user(nick, password,birthday,address,user ) values(?,?,?,?,?)",user.Nick,user.Password,user.Birthday,user.Address,user.User)
	if err != nil {
		//保存数据时遇到错误
		return -1,err
	}

	row, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}
	return row,err

}


//查询用户
func QueryUser()  {
	Db.QueryRow("select *from ")


}