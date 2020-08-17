package handlers

import (
	"fmt"
	"whisper/serves"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

//Login handler 处理登录请求
func Login(ctx iris.Context) {
	var res structs.ResLogin
	err := ctx.ReadJSON(&res)
	if err != nil {
		fmt.Println("前端传入数据出错", err.Error())
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

//Regist handler 处理注册请求
func Regist(ctx iris.Context) {
	var res structs.ResRegist
	err := ctx.ReadJSON(&res)
	if err != nil {
		fmt.Println("前端传入数据出错", err.Error())
		ctx.StatusCode(400)
		ctx.WriteString("传入格式不正确" + err.Error())
		return
	}
}
