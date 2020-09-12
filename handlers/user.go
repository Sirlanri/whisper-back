package handlers

import (
	"fmt"
	"whisper/serves"
	"whisper/sqls"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

/*GetUserInfo handler
登录后，获取当前登录用户的信息  */
func GetUserInfo(ctx iris.Context) {
	userid := serves.GetUserID(ctx)
	//mail := ctx.URLParams().Get("mail")
	result := sqls.GetUserInfo(userid)
	ctx.JSON(result)
}

/*ChangeAvatar handler
接收新头像的URL */
func ChangeAvatar(ctx iris.Context) {
	//获取用户id
	userid := serves.GetUserID(ctx)
	//新头像的URL
	avatar := ctx.URLParam("url")
	if avatar == "" {
		ctx.StatusCode(404)
		ctx.WriteString("传入参数有误")
		return
	}
	result := sqls.ChangeAvatar(avatar, userid)
	if !result {
		ctx.StatusCode(404)
		ctx.WriteString("修改头像URL出错")
	} else {
		ctx.WriteString("修改头像成功")
	}
}

/*ChangeBannar handler
接收新bannar的URL */
func ChangeBannar(ctx iris.Context) {
	//获取用户id
	userid := serves.GetUserID(ctx)
	//新头像的URL
	bannar := ctx.URLParam("url")
	if bannar == "" {
		ctx.StatusCode(404)
		ctx.WriteString("传入参数有误")
		return
	}
	result := sqls.ChangeBannar(bannar, userid)
	if !result {
		ctx.StatusCode(404)
		ctx.WriteString("修改bannar URL出错")
	} else {
		ctx.WriteString("修改bannar成功")
	}
}

/*ChangeInfo handler
用户修改资料，传入昵称、邮箱、个人介绍*/
func ChangeInfo(ctx iris.Context) {
	//获取用户id
	userid := serves.GetUserID(ctx)
	var res structs.ResChangeInfo
	err := ctx.ReadJSON(&res)
	if err != nil {
		fmt.Println("修改资料，前端传入数据有误", err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("前端传入数据有误")
		return
	}
	//检测用户名和邮箱是否合法
	valuable := serves.Check(res.Name, res.Mail)
	if !valuable {
		ctx.StatusCode(iris.StatusNotAcceptable)
		ctx.WriteString("用户名或邮箱不合法")
		return
	}

	//传入SQL部分
	result := sqls.ChangeInfo(res, userid)
	if !result {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("写入用户资料出错")
		return
	}
	ctx.WriteString("修改资料成功，请刷新页面查看")

}
