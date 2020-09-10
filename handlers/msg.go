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

//ReadMsg handler 将某条消息标为已读
func ReadMsg(ctx iris.Context) {
	//从session中获取用户邮箱
	mail := serves.GetUserMail(ctx)
	replyid, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("更改状态失败，传入参数有误")
	}
	//执行SQL 更改已读状态
	result := sqls.ReadMsg(mail, replyid)
	if !result {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("标记已读失败")

	}
}
