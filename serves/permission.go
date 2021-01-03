package serves

import (
	"fmt"
	"whisper/sqls"

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
	sessionUser.Set("userid", sqls.GetIDBymail(mail))
}

//AdminPermiss 管理员登录后，设置管理员权限
func AdminPermiss(ctx iris.Context, mail string) {
	sessionAdmin := sess.Start(ctx)
	sessionAdmin.Set("admin", true)
	sessionAdmin.Set("isLogin", true)
	sessionAdmin.Set("mail", mail)
	sessionAdmin.Set("userid", sqls.GetIDBymail(mail))
}

//GetUserMail 从Session中获取用户邮箱
func GetUserMail(ctx iris.Context) (mail string) {
	sessionMail := sess.Start(ctx)
	mail = sessionMail.GetString("mail")
	return
}

//GetUserID 从session中获取用户的id
func GetUserID(ctx iris.Context) int {
	sessionID := sess.Start(ctx)
	//如果cookie有问题，没有这个key，就返回0
	userid, err := sessionID.GetInt("userid")
	if err != nil {
		userid = 0
	}
	return userid
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

/*IsAdmin 当前登录用户是否为管理员*/
func IsAdmin(ctx iris.Context) bool {
	sessAdmin := sess.Start(ctx)
	check := sessAdmin.GetBooleanDefault("admin", false)
	return check
}
