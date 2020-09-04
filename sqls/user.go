package sqls

import (
	"fmt"
	"whisper/structs"
)

/*GetUserInfo SQL 获取用户的信息
传入mail，ctx */
func GetUserInfo(mail string) (result structs.UserInfo) {
	tx, _ := Db.Begin()
	var userid int
	row1 := tx.QueryRow("select userid,userName,intro,avatar,bannar from user where mail=?", mail)
	row1.Scan(&userid, &result.Name, &result.Intro, &result.Avatar, &result.Bannar)

	counts, err := tx.Query(`select count(*) from post where publisher=? 
	union select count(*) from reply where fromUser=?`, userid, userid)
	if err != nil {
		println("query出错", err.Error())
	}
	/*
		counts.Next()
		err = counts.Scan(&result.PostCount)
		if err != nil {
			println("post写入出错", err.Error())
		}
		//counts.Next()
		err = counts.Scan(&result.ReplyCount)
		if err != nil {
			println("reply写入出错", err.Error())
		}
	*/
	index := true
	for counts.Next() {
		if index {
			err = counts.Scan(&result.PostCount)
			if err != nil {
				println("post写入出错", err.Error())
			}
			index = false
		} else {
			counts.Next()
			err = counts.Scan(&result.ReplyCount)
			if err != nil {
				println("reply写入出错", err.Error())
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("获取用户信息-执行SQL出错 ", err.Error())
	}
	return
}
