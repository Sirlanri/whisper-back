package handlers

import "github.com/kataras/iris/v12"

/*GetUserInfo handler
获取用户的信息  */
func GetUserInfo(ctx iris.Context) {
	mail := ctx.Params().Get("mail")

}
