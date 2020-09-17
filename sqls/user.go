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
	err = row1.Scan(&result.Mail, &result.Name, &result.Intro, &result.Avatar, &result.Bannar, &powerid)
	if err != nil {
		fmt.Println("读取用户信息失败", err.Error())
	}
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

/*GetUserInfoByName SQL
通过name获取用户的信息
传入用户的name */
func GetUserInfoByName(userName string) (result structs.UserInfo) {
	tx, _ := Db.Begin()
	var (
		err     error
		powerid int
		userid  int
	)
	//通过name获取id
	useridRow := tx.QueryRow(`select userid from user where userName=?`, userName)
	err = useridRow.Scan(&userid)
	if err != nil {
		fmt.Println("通过name获取id出错", err.Error())
		return
	}

	row1 := tx.QueryRow("select mail,userName,intro,avatar,bannar,power from user where userid=?", userid)
	err = row1.Scan(&result.Mail, &result.Name, &result.Intro, &result.Avatar, &result.Bannar, &powerid)
	if err != nil {
		fmt.Println("读取用户信息失败", err.Error())
	}
	if powerid == 1 {
		result.Power = "user"
	} else {
		result.Power = "admin"
	}
	postcount := tx.QueryRow("select count(*) from post where publisher=?", userid)
	err = postcount.Scan(&result.PostCount)
	if err != nil {
		fmt.Println("post写入出错", err.Error())
	}

	replyCount := tx.QueryRow("select count(*) from reply where fromUser=?", userid)
	err = replyCount.Scan(&result.ReplyCount)
	if err != nil {
		fmt.Println("reply写入出错", err.Error())
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

/*ChangeInfo SQL
修改用户资料  传入修改结构体和id*/
func ChangeInfo(res structs.ResChangeInfo, userid int) bool {
	tx, _ := Db.Begin()
	_, err := tx.Exec(`update user set mail=?,userName=?,intro=? 
	where userid=?`, res.Mail, res.Name, res.Intro, userid)
	if err != nil {
		fmt.Println("修改用户资料失败", err.Error())
		return false
	}
	tx.Commit()
	return true
}

/*DelUserByPostID SQL
删除某个用户及其post，通过post的id
*/
func DelUserByPostID(postid int) bool {
	tx, _ := Db.Begin()

	//获取用户的id
	useridRow := tx.QueryRow(`select publisher from post where postid=?`, postid)
	var userid int
	err := useridRow.Scan(&userid)
	if err != nil {
		fmt.Println("通过post获取用户id失败", err.Error())
		return false
	}

	//删除用户
	_, err = tx.Exec(`delete from user where userid=?`, userid)
	if err != nil {
		fmt.Println("删除用户失败", err.Error())
		return false
	}

	//删除用户发布的post
	_, err = tx.Exec(`delete from post where publisher=?`, userid)
	if err != nil {
		fmt.Println("删除用户发布的post出错", err.Error())
		return false
	}
	return true

}
