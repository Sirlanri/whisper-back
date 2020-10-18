package handlers

import (
	"fmt"
	"whisper/serves"
	"whisper/sqls"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

//GetAllReply handler 获取全部回复详情 采用懒加载，此handler暂时弃用
func GetAllReply(ctx iris.Context) {
	//从session中获取用户id
	userid := serves.GetUserID(ctx)

	//返回0 表示未登录
	if userid == 0 {
		ctx.StatusCode(210)
		ctx.WriteString("未登录账号")
		return
	}

	replys := sqls.GetAllReply(userid)
	jsonData := map[string][]structs.ReplyDetail{
		"replys": replys,
	}
	ctx.JSON(jsonData)
}

/*GetReplys handler 用以取代GetAllReply
获取回复，传入num，一次返回20个*/
func GetReplys(ctx iris.Context) {
	//从session中获取用户id
	userid := serves.GetUserID(ctx)

	//返回0 表示未登录
	if userid == 0 {
		ctx.StatusCode(210)
		ctx.WriteString("未登录账号")
		return
	}
	//获取请求的起始值
	num := ctx.URLParamIntDefault("num", 0)
	replys := sqls.GetReplys(userid, num)
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

	//检测用户是否已登录
	userid := serves.GetUserID(ctx)
	if userid == 0 {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("用户未登录")
		return
	}

	result, info := sqls.NewReply(resRelpy, userid)
	if !result {
		ctx.StatusCode(iris.StatusUnauthorized)
	}
	ctx.WriteString(info)
}

/*DelReply handler 删除某条回复
传入postid 需验证用户是否为接收人*/
func DelReply(ctx iris.Context) {
	replyid, err := ctx.URLParamInt("id")
	if err != nil {
		fmt.Println("删除reply，传入参数错误", err.Error())
	}

	//检测用户是否已登录+获取用户id
	userid := serves.GetUserID(ctx)
	if userid == 0 {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("用户未登录")
		return
	}
	//执行删除SQL
	result := sqls.DelReply(replyid, userid)
	if !result {
		ctx.StatusCode(210)
		ctx.WriteString("删除回复失败")
		return
	}
	ctx.WriteString("删除回复成功")
}
