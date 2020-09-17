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
		ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("传入格式有误")
		return
	}
	if ResPost.Content == "" {
		ctx.StatusCode(403)
		ctx.WriteString("内容不能为空哦")
		return
	}
	//从session中获取用户mail
	userid := serves.GetUserID(ctx)
	if userid == 0 {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("用户未登录")
		return
	}
	sqls.NewPost(ResPost, userid)
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
	posts := sqls.GetALlPost()
	jsondata := map[string][]structs.DataPost{
		"posts": posts,
	}
	ctx.JSON(jsondata)
}

/*GetPostByUser handler
通过用户名获取某个用户的post*/
func GetPostByUser(ctx iris.Context) {
	name := ctx.URLParam("name")
	posts := sqls.GetALlPostByUser(name)
	jsondata := map[string][]structs.DataPost{
		"posts": posts,
	}
	ctx.JSON(jsondata)
}

/*GetPostByGroup handler
通过群组id，获取此群组的post 限制20*/
func GetPostByGroup(ctx iris.Context) {
	groupid, err := ctx.URLParamInt("id")
	if err != nil {
		fmt.Println("传入参数错误", err.Error())
		ctx.StatusCode(404)
		ctx.WriteString("传入参数错误")
		return
	}
	posts := sqls.GetPostByGroup(groupid)
	jsondata := map[string][]structs.DataPost{
		"posts": posts,
	}
	ctx.JSON(jsondata)
}

/*DelPost handler
删除post 传入post的id*/
func DelPost(ctx iris.Context) {
	postid, err := ctx.URLParamInt("id")
	if err != nil {
		fmt.Println("前端传入数据不合法", err.Error())
		ctx.StatusCode(404)
		ctx.WriteString("传入数据不合法")
		return
	}
	result := sqls.DelPost(postid)
	if !result {
		ctx.StatusCode(404)
		ctx.WriteString("删除Post失败")
	} else {
		ctx.WriteString("删除成功")
	}
}

/*DelMyPost handler 删除自己发送的某条post
传入postid*/
func DelMyPost(ctx iris.Context) {
	postid, err := ctx.URLParamInt("id")
	if err != nil {
		fmt.Println("删除自己发送的post，传入值有误", err.Error())
		ctx.StatusCode(404)
		ctx.WriteString("传入参数有误")
	}
	userid := serves.GetUserID(ctx)
	result := sqls.DelMyPost(postid, userid)
	if !result {
		ctx.StatusCode(404)
		ctx.WriteString("传入参数有误")
		return
	}

}
