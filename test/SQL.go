package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Db 创建的唯一指针
var Db *sql.DB

//初始化，自动创建db指针
func main() {
	Db = ConnectDB()
	insertalot()
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

//Test 测试的
func Test() {
	tx, err := Db.Begin()
	if err != nil {
		println(err.Error())
	}
	t1 := time.Now()

	for i := 0; i < 10000; i++ {
		_, err := tx.Exec("insert into `test` (`test`) values (?)", i)
		if err != nil {
			println("执行出错", err.Error())
		}
	}

	elapsed := time.Since(t1)
	err = tx.Commit()
	println("用时", elapsed.Seconds())

	if err != nil {
		println(err.Error())
	}

}

func insertalot() {
	tx, _ := Db.Begin()

	//计时
	t1 := time.Now()
	for i := 0; i < 10000; i++ {
		word := "当前插入数据："
		tx.Exec(`insert into post (publisher,groupid,content) 
		values (?,?,?)`, 1, 24, word)

	}

	tx.Commit()
	elapsed := time.Since(t1)
	fmt.Println("共计耗时 ", elapsed)
}

func updateTest() {
	tx, _ := Db.Begin()

	//计时
	t1 := time.Now()
	for i := 1000; i < 1003; i++ {
		word := "当前插入数据：" + string(i)
		_, err := tx.Exec(`update post set content=? where postid=?`, word, i)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	tx.Commit()
	elapsed := time.Since(t1)
	fmt.Println("共计耗时 ", elapsed)
}
