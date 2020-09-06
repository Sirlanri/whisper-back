package serves

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

//保存用户登录状态信息
var (
	cookieName = "login"
	sess       = sessions.New(sessions.Config{Cookie: cookieName})
)

//VisitorPermiss 用户登录成功后，设置对应的权限
func VisitorPermiss(ctx iris.Context, mail string) {
	sessionUser := sess.Start(ctx)
	sessionUser.Set("isLogin", true)
	sessionUser.Set("mail", mail)
}

//AdminPermiss 管理员登录后，设置管理员权限
func AdminPermiss(ctx iris.Context, mail string) {
	sessionAdmin := sess.Start(ctx)
	sessionAdmin.Set("admin", true)
	sessionAdmin.Set("isLogin", true)
	sessionAdmin.Set("mail", mail)
}

//GetUserMail 从Session中获取用户邮箱
func GetUserMail(ctx iris.Context) (mail string) {
	sessionMail := sess.Start(ctx)
	mail = sessionMail.GetString("mail")
	return
}

//ClearPermiss 注销登录，清除用户权限
func ClearPermiss(ctx iris.Context) {
	sessClear := sess.Start(ctx)
	mail := sessClear.GetString("mail")
	sessClear.Destroy()
	ctx.RemoveCookie("login")
	fmt.Println("用户注销 ", mail)
	ctx.WriteString("注销成功")
}
