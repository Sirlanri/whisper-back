package sqls

import (
	"fmt"
	"time"
	"whisper/structs"
)

//NewPost 负责处理前端接收到的数据
func NewPost(res structs.ResPost, mail string) {
	var (
		err             error
		userid, groupid int
	)
	//计时
	t1 := time.Now()

	tx, _ := Db.Begin()
	//通过邮箱获取用户id
	idRow := tx.QueryRow("select userid from user where mail=?", mail)
	err = idRow.Scan(&userid)
	if err != nil {
		fmt.Println("SQL获取用户ID出错", err.Error())
	}
	fmt.Println("获取用户ID ", time.Since(t1))

	//获取群组id
	groupRow := tx.QueryRow("select groupid from `group` where groupName=?", res.Group)
	err = groupRow.Scan(&groupid)
	if err != nil {
		fmt.Println("获取群ID出错", err.Error())
	}
	fmt.Println("获取群组ID ", time.Since(t1))

	//写入post表
	pre, err := tx.Prepare(`INSERT INTO post (publisher,groupid,content)
	 values (?,?,?)`)
	if err != nil {
		fmt.Println("预编译SQL出错", err.Error())
	}
	newrow, err := pre.Exec(userid, groupid, res.Content)
	if err != nil {
		fmt.Println("SQL写入post出错 ", err.Error())
	}
	fmt.Println("写入post ", time.Since(t1))
	postid, _ := newrow.LastInsertId()

	//插入picture表
	for _, pic := range res.Pics {
		_, err = tx.Exec("insert into picture values (?,?)", postid, pic)
		if err != nil {
			fmt.Println("SQL 插入图片出错", err.Error())
		}
	}
	fmt.Println("写入图片 ", time.Since(t1))

	//写入tag
	for _, tag := range res.Tags {
		_, err = tx.Exec("insert into tag values (?,?)", tag, postid)
		if err != nil {
			fmt.Println("SQL 插入Tag出错", err.Error())
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("SQL commit出错 ", err.Error())
	}
	elapsed := time.Since(t1)
	fmt.Println("共计耗时 ", elapsed)
	fmt.Println("SQL写入完毕")
}

func NewPost2(res structs.ResPost, mail string) {
	var (
		err             error
		userid, groupid int
	)
	//计时
	t1 := time.Now()

	//通过邮箱获取用户id
	idRow := Db.QueryRow("select userid from user where mail=?", mail)
	err = idRow.Scan(&userid)
	if err != nil {
		fmt.Println("SQL获取用户ID出错", err.Error())
	}
	fmt.Println("获取用户ID ", time.Since(t1))

	//获取群组id
	groupRow := Db.QueryRow("select groupid from `group` where groupName=?", res.Group)
	err = groupRow.Scan(&groupid)
	if err != nil {
		fmt.Println("获取群ID出错", err.Error())
	}
	fmt.Println("获取群组ID ", time.Since(t1))

	//写入post表
	pre, err := Db.Prepare(`INSERT INTO post (publisher,groupid,content)
	 values (?,?,?)`)
	if err != nil {
		fmt.Println("预编译SQL出错", err.Error())
	}
	newrow, err := pre.Exec(userid, groupid, res.Content)
	if err != nil {
		fmt.Println("SQL写入post出错 ", err.Error())
	}
	fmt.Println("写入post ", time.Since(t1))
	postid, _ := newrow.LastInsertId()

	//插入picture表
	for _, pic := range res.Pics {
		_, err = Db.Exec("insert into picture values (?,?)", postid, pic)
		if err != nil {
			fmt.Println("SQL 插入图片出错", err.Error())
		}
	}
	fmt.Println("写入图片 ", time.Since(t1))

	if err != nil {
		fmt.Println("SQL commit出错 ", err.Error())
	}
	elapsed := time.Since(t1)
	fmt.Println("共计耗时 ", elapsed)
	fmt.Println("SQL写入完毕")
}
