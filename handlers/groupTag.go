package handlers

import (
	"fmt"
	"whisper/serves"
	"whisper/sqls"
	"whisper/structs"

	"github.com/kataras/iris/v12"
)

//GetGroups handlers 获取群组列表
func GetGroups(ctx iris.Context) {
	groups := sqls.Getgroups()
	jsondata := map[string][]structs.GroupFront{
		"groups": groups,
	}
	ctx.JSON(jsondata)
}

//NewGroup handler 创建一个群
func NewGroup(ctx iris.Context) {
	var res structs.ResGroup
	ctx.ReadJSON(&res)

	//检测用户是否已登录
	userid := serves.GetUserID(ctx)
	if userid == 0 {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.WriteString("用户未登录")
		return
	}

	//检测三个值是否都存在
	if res.Intro != "" && res.Name != "" && res.Pic != "" {
		result := sqls.NewGroup(res)
		if result {
			fmt.Println("成功创建群组：", res.Name)
			ctx.WriteString("创建成功")
		} else {
			ctx.WriteString("创建失败，请检查此群名是否已存在")
		}
	} else {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("前端传入值不正确")
		return
	}

}

/*DelGroupOnly handler
删除群组信息，把群内的post修改为无群状态*/
func DelGroupOnly(ctx iris.Context) {
	groupid, err := ctx.URLParamInt("id")
	if err != nil {
		fmt.Println("删除群，传入id错误")
		ctx.StatusCode(404)
		ctx.WriteString("传入id不合法")
		return
	}
	result := sqls.DelGroupOnly(groupid)
	if !result {
		ctx.StatusCode(404)
		ctx.WriteString("删除群失败")
	} else {
		ctx.WriteString("删除成功")
	}
}

/*DelGroupAll handler
删除群信息，并删除群内全部post*/
func DelGroupAll(ctx iris.Context) {
	groupid, err := ctx.URLParamInt("id")
	if err != nil {
		fmt.Println("删除群，传入id错误")
		ctx.StatusCode(404)
		ctx.WriteString("传入id不合法")
		return
	}
	result := sqls.DelGroupAll(groupid)
	if !result {
		ctx.StatusCode(404)
		ctx.WriteString("删除群失败")
	} else {
		ctx.WriteString("删除成功")
	}
}
