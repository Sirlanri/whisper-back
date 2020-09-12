package handlers

import (
	"fmt"
	"whisper/serves"
	"whisper/sqls"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

/*Login handler 处理登录请求

返回code：200-普通用户登录 201-管理员登录 202-用户名或密码错误*/
//400-前端传入json格式不正确
func Login(ctx iris.Context) {
	var res structs.ResLogin
	err := ctx.ReadJSON(&res)
	if err != nil {
		fmt.Println("前端传入数据出错", err.Error())
		ctx.StatusCode(400)
		ctx.WriteString("传入格式不正确" + err.Error())
		return
	}
	result, power := serves.Login(res.Mail, res.Password)
	if result && power == 0 {
		serves.AdminPermiss(ctx, res.Mail)
		ctx.StatusCode(201)
		ctx.WriteString("欢迎管理员登录")
		return
	}
	if result {
		serves.VisitorPermiss(ctx, res.Mail)
		ctx.WriteString("登录成功")
	} else {
		ctx.StatusCode(202)
		ctx.WriteString("用户名或密码错误")
	}
	return
}

/*Regist handler 处理注册请求

返回code：200-注册成功 201-用户名密码已存在 202-用户名邮箱格式不正确
400-前端传入json格式不正确*/
func Regist(ctx iris.Context) {
	var res structs.ResRegist
	err := ctx.ReadJSON(&res)
	if err != nil {
		fmt.Println("前端传入数据出错", err.Error())
		ctx.StatusCode(400)
		ctx.WriteString("传入格式不正确" + err.Error())
		return
	}
	result, code := serves.Regist(res.Name, res.Mail, res.Password)
	ctx.StatusCode(code)
	ctx.WriteString(result)
}

/*Logout handler 注销登录
 */
func Logout(ctx iris.Context) {
	serves.ClearPermiss(ctx)
}

/*GetUserInfoByCookie handler
通过cookie获取用户信息，用于刷新页面后免登录*/
func GetUserInfoByCookie(ctx iris.Context) {
	userid := serves.GetUserID(ctx)
	if userid == 0 {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("cookie无效，请重新登录")
		return
	}
	result := sqls.GetUserInfo(userid)
	ctx.JSON(result)
}
