package handlers

import (
	"fmt"
	"whisper/sqls"

	"github.com/kataras/iris/v12"
)

/*GetUserInfo handler
获取用户的信息  */
func GetUserInfo(ctx iris.Context) {
	mail := ctx.URLParam("mail")
	//mail := ctx.URLParams().Get("mail")
	result := sqls.GetUserInfo(mail)
	fmt.Println(result)
}
