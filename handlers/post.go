package handlers

import (
	"fmt"
	"io"
	"os"
	"whisper/serves"
	"whisper/sqls"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

//NewPost handler 发布新推文
func NewPost(ctx iris.Context) {
	var ResPost structs.ResPost
	err := ctx.ReadJSON(&ResPost)

	if err != nil {
		fmt.Println("NewPost出错，前端传入格式错误", err.Error())
		ctx.WriteString("传入格式有误")
	}
	//从session中获取用户mail
	mail := serves.GetUserMail(ctx)
	if mail == "" {
		ctx.WriteString("用户未登录")
		return
	}
	sqls.NewPost(ResPost, mail)
}

/*UploadPics handler 上传图片
并命名为UUID保存到uploadpics目录下，向前端返回完整的URL*/
func UploadPics(ctx iris.Context) {
	file, info, err := ctx.FormFile("img")
	if err != nil {
		//status==500 上传失败
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("上传图片失败")
		println("上传图片失败", err.Error())
		return
	}
	defer file.Close()
	fname := serves.Createid() + info.Filename
	//图片保存目录 uploadpics
	out, err := os.OpenFile("./uploadpics/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString("图片保存至服务器失败")
		println("图片保存至服务器失败", err.Error())
		return
	}
	defer out.Close()
	io.Copy(out, file)
	whole := "http://localhost:8090/whisper/getpics/" + fname
	ctx.WriteString(whole)
	fmt.Println("上传图片成功", fname)
}

//GetGroupNames handler 创建post时，获取全部群组列表
func GetGroupNames(ctx iris.Context) {
	result := sqls.GetGroupNames()
	jsondata := map[string][]string{
		"groups": result,
	}
	ctx.JSON(jsondata)
}

//GetTags handler 创建post时，获取全部群组列表
func GetTags(ctx iris.Context) {
	result := sqls.GetTags()
	jsondata := map[string][]string{
		"tags": result,
	}
	ctx.JSON(jsondata)
}

//GetAllPost handler 获取全部的post
func GetAllPost(ctx iris.Context) {
	posts := sqls.GetALlPost2()
	jsondata := map[string][]structs.DataPost{
		"posts": posts,
	}
	ctx.JSON(jsondata)
}
