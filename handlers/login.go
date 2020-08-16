package handlers

import (
	"whisper/serves"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

//Login handler 处理登录请求
func Login(ctx iris.Context) {
	var res structs.ResLogin
	err := ctx.ReadJSON(&res)
	if err != nil {
		println("前端传入数据出错", err.Error())
		ctx.WriteString("传入格式不正确" + err.Error())
		return
	}
	result := serves.Login(res.Mail, res.Password)
	if result {
		ctx.WriteString("登录成功")
	} else {
		ctx.WriteString("用户名或密码错误")
		ctx.StatusCode(iris.StatusUnauthorized)
	}
}
