package handlers

import (
	"whisper/serves"
	"whisper/sqls"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

//GetAllReply handler 获取全部回复详情
func GetAllReply(ctx iris.Context) {
	//从session中获取用户邮箱
	mail := serves.GetUserMail(ctx)
	replys := sqls.GetAllReply(mail)
	jsonData := map[string][]structs.ReplyDetail{
		"replys": replys,
	}
	ctx.JSON(jsonData)
}
