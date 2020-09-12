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
