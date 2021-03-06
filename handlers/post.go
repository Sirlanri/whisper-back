package handlers

import (
	"fmt"
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
	ctx.Redirect("http://localhost:8091/img/upload")

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
	nums := ctx.URLParamIntDefault("num", 0)
	posts := sqls.GetAllPostByUser(name, nums)
	jsondata := map[string][]structs.DataPost{
		"posts": posts,
	}
	ctx.JSON(jsondata)
}

/*GetPostByGroup handler 懒加载
通过群组id，获取此群组的post 限制20*/
func GetPostByGroup(ctx iris.Context) {
	groupid, err := ctx.URLParamInt("id")
	if err != nil {
		fmt.Println("传入参数错误", err.Error())
		ctx.StatusCode(404)
		ctx.WriteString("传入参数错误")
		return
	}
	num := ctx.URLParamIntDefault("num", 0)
	posts := sqls.GetPostByGroup(groupid, num)
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

/*GetLazyPost handler
懒加载的方式获取post，传入一个值n，返回[n,n+20)条*/
func GetLazyPost(ctx iris.Context) {
	num := ctx.URLParamIntDefault("num", 0)
	posts := sqls.GetLazyPost(num)
	jsondata := map[string][]structs.DataPost{
		"posts": posts,
	}
	ctx.JSON(jsondata)
}
