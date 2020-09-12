package handlers

import (
	"whisper/serves"
	"whisper/sqls"

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
