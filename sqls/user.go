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
	postCount := tx.QueryRow("select count(*) from post where publisher=?", userid)
	replyCount := tx.QueryRow("select count(*) from reply where publisher=?", userid)
	postCount.Scan(&result.PostCount)
	replyCount.Scan(&result.ReplyCount)
	err := tx.Commit()
	if err != nil {
		fmt.Println("获取用户信息-执行SQL出错 ", err.Error())
	}
	return
}
