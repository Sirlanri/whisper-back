package sqls

import (
	"fmt"
	"whisper/structs"
)

//GetAllReply SQL 获取全部的回复
func GetAllReply(userid int) (replys []structs.ReplyDetail) {
	tx, _ := Db.Begin()

	//通过用户ID查询此用户接收的reply
	replyRow, err := tx.Query(`select replyid,postid,fromUser,content,haveRead 
	from reply where toUser=? ORDER BY replyid DESC`, userid)
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
		err = replyRow.Scan(&reply.ID, &reply.Postid, &userid, &reply.Content, &reply.HaveRead)
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
func ReadMsg(userid int, replyid int) bool {
	tx, _ := Db.Begin()

	//通过用户ID
	_, err := tx.Exec(`UPDATE reply SET haveRead=1 WHERE replyid=? AND toUser=?`, replyid, userid)
	if err != nil {
		fmt.Println("更改reply已读状态失败", err.Error())
		return false
	}
	tx.Commit()
	return true
}

//NewReply SQL 将新回复插入数据库
func NewReply(reply structs.ResNewReply, userid int) (result bool, info string) {
	tx, _ := Db.Begin()

	result = true
	//通过被回复人的name获取其id
	useridRow := tx.QueryRow(`select userid from user where userName=?`, reply.Name)
	var (
		receverid int
	)

	//被回复人的id
	err := useridRow.Scan(&receverid)
	if err != nil {
		fmt.Println("找不到此用户：", err.Error())
		info = "找不到此用户"
		result = false
	}

	//将回复写入数据库
	_, err = tx.Exec(`insert into reply (postid,fromUser,toUser,content)
   values (?,?,?,?)`, reply.ID, userid, receverid, reply.Content)
	if err != nil {
		fmt.Println("reply写入数据库出错", err.Error())
		info = "reply写入数据库出错"
		result = false
	}
	tx.Commit()
	info = "回复成功"

	return
}

/*DelReply SQL 删除回复
传入replyid userid*/
func DelReply(replyid, userid int) bool {
	tx, _ := Db.Begin()

	_, err := tx.Exec(`delete from reply where replyid=? and toUser=?`,
		replyid, userid)
	if err != nil {
		fmt.Println("SQL 删除reply回复失败", err.Error())
		return false
	}
	tx.Commit()
	return true
}

//GetReplys SQL 懒加载回复，一次20条
func GetReplys(userid, num int) (replys []structs.ReplyDetail) {
	tx, _ := Db.Begin()

	//通过用户ID查询此用户接收的reply
	replyRow, err := tx.Query(`select replyid,postid,fromUser,content,haveRead 
  from reply where toUser=? ORDER BY replyid DESC LIMIT ?,20`, userid, num)
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
		err = replyRow.Scan(&reply.ID, &reply.Postid, &userid, &reply.Content, &reply.HaveRead)
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
