package handlers

import (
	"fmt"
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
