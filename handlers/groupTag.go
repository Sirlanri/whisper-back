package handlers

import (
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
