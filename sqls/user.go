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
	var err error
	row1 := tx.QueryRow("select userid,userName,intro,avatar,bannar from user where mail=?", mail)
	row1.Scan(&userid, &result.Name, &result.Intro, &result.Avatar, &result.Bannar)

	postcount := tx.QueryRow("select count(*) from post where publisher=?", userid)
	err = postcount.Scan(&result.PostCount)
	if err != nil {
		println("post写入出错", err.Error())
	}

	replyCount := tx.QueryRow("select count(*) from reply where fromUser=?", userid)
	err = replyCount.Scan(&result.ReplyCount)
	if err != nil {
		println("reply写入出错", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("获取用户信息-执行SQL出错 ", err.Error())
	}
	return
}
