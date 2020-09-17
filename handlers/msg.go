package handlers

import (
	"whisper/serves"
	"whisper/sqls"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

//GetAllReply handler 获取全部回复详情
func GetAllReply(ctx iris.Context) {
	//从session中获取用户id
	userid := serves.GetUserID(ctx)

	replys := sqls.GetAllReply(userid)
	jsonData := map[string][]structs.ReplyDetail{
		"replys": replys,
	}
	ctx.JSON(jsonData)
}

//ReadMsg handler 将某条消息标为已读
func ReadMsg(ctx iris.Context) {
	//从session中获取用户邮箱
	userid := serves.GetUserID(ctx)
	replyid, err := ctx.URLParamInt("id")
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("更改状态失败，传入参数有误")
	}
	//执行SQL 更改已读状态
	result := sqls.ReadMsg(userid, replyid)
	if !result {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("标记已读失败")

	}
}

//NewReply handler 新回复
func NewReply(ctx iris.Context) {
	var resRelpy structs.ResNewReply
	ctx.ReadJSON(&resRelpy)
	if resRelpy.Content == "" {
		ctx.StatusCode(403)
		ctx.WriteString("内容不能为空哦")
		return
	}
	userid := serves.GetUserID(ctx)
	result, info := sqls.NewReply(resRelpy, userid)
	if !result {
		ctx.StatusCode(iris.StatusUnauthorized)
	}
	ctx.WriteString(info)
}
