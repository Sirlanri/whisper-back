package sqls

import (
	"fmt"
	"whisper/structs"
)

/*GetUserInfo SQL 获取用户的信息
传入用户的userid */
func GetUserInfo(userid int) (result structs.UserInfo) {
	tx, _ := Db.Begin()
	var (
		err     error
		powerid int
	)
	row1 := tx.QueryRow("select mail,userName,intro,avatar,bannar,power from user where userid=?", userid)
	row1.Scan(&result.Mail, &result.Name, &result.Intro, &result.Avatar, &result.Bannar, &powerid)
	if powerid == 1 {
		result.Power = "user"
	} else {
		result.Power = "admin"
	}
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

/*ChangeAvatar SQL
改变用户头像，将新URL写入数据库
操作成功，返回true*/
func ChangeAvatar(url string, userid int) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`update user set avatar=? where userid=?`, url, userid)
	if err != nil {
		fmt.Println("升级头像URL出错", err.Error())
		return false
	}
	tx.Commit()
	return true
}

/*ChangeBannar SQL
改变用户头像，将新URL写入数据库
操作成功，返回true*/
func ChangeBannar(url string, userid int) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`update user set bannar=? where userid=?`, url, userid)
	if err != nil {
		fmt.Println("升级bannar URL出错", err.Error())
		return false
	}
	tx.Commit()
	return true
}
