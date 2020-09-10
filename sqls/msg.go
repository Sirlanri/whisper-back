package sqls

import (
	"fmt"
	"whisper/structs"
)

//GetAllReply SQL 获取全部的回复
func GetAllReply(mail string) (replys []structs.ReplyDetail) {
	tx, _ := Db.Begin()

	//通过mail获取当前登录用户的id
	var userid int
	idRow := tx.QueryRow(`select userid from user where mail=?`, mail)
	idRow.Scan(&userid)

	//通过用户ID查询此用户接收的reply
	replyRow, err := tx.Query(`select replyid,fromUser,content,haveRead 
	from reply where toUser=?`, userid)
	if err != nil {
		fmt.Println("查询用户回复出错", err.Error())
	}
	var (
		userids []int //消息发送者的ID列表
	)

	//写入id,content,haveRead属性
	for replyRow.Next() {
		var (
			reply  structs.ReplyDetail //暂存回复
			userid int
		)
		err = replyRow.Scan(&reply.ID, &userid, &reply.Content, &reply.HaveRead)
		if err != nil {
			fmt.Println("写入ReplyRow出错", err.Error())
		}
		userids = append(userids, userid)
		replys = append(replys, reply)
	}

	//同过用户id获取用户的name和avatar
	for index, userid := range userids {
		userRow := tx.QueryRow(`select userName,avatar from user where userid=?`, userid)
		userRow.Scan(&replys[index].Name, &replys[index].Avatar)
	}
	return
}

//ReadMsg SQL 验证权限，将reply标为已读
func ReadMsg(mail string, replyid int) bool {
	tx, _ := Db.Begin()
	//通过mail获取用户id
	idrow := tx.QueryRow(`select userid from user where mail=?`, mail)
	var userid int
	idrow.Scan(&userid)

	//通过用户ID
	_, err := tx.Exec(`UPDATE reply SET haveRead=1 WHERE replyid=? AND toUser=?`, replyid, userid)
	if err != nil {
		fmt.Println("更改reply已读状态失败", err.Error())
		return false
	}
	tx.Commit()
	return true
}
