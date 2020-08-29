package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//Db 创建的唯一指针
var Db *sql.DB

//初始化，自动创建db指针
func main1() {
	Db = ConnectDB()
	Regist("蓝", "ii@ri.cn", "123")
}

//ConnectDB 初始化时，连接数据库
func ConnectDB() *sql.DB {
	Db, err := sql.Open("mysql", "root:123456@/whisper")
	if err != nil {
		fmt.Println("数据库初始化链接失败", err.Error())
	}

	if Db.Ping() != nil {
		fmt.Println("初始化-数据库-用户/密码/库验证失败", Db.Ping().Error())
		return nil
	}
	return Db
}

//Regist 注册函数
func Regist(name, mail, pw string) (result string) {
	//传入的参数分别对应着mail userName password userName mail
	pre, err := Db.Prepare(`INSERT INTO user (mail, userName, password)
		SELECT ?,?,?
		from DUAL
		WHERE not exists (
				SELECT *
				from user
				WHERE userName = ? or mail = ? LIMIT 1);`)
	if err != nil {
		fmt.Println("预编译表达式出错", err.Error())
	}
	effects, err := pre.Exec(mail, name, pw, name, mail)
	if err != nil {
		fmt.Println("写入用户数据，执行SQL出错", err.Error())
	}
	//如果rownum==0，说明没有插入数据，即用户名邮箱已存在
	rownum, _ := effects.RowsAffected()
	println(rownum)
	result = "执行完毕"
	return
}
