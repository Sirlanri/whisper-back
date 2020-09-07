package sqls

import (
	"fmt"
	"whisper/structs"
)

//Getgroups SQL 获取群组列表
func Getgroups() []structs.GroupFront {
	var group structs.GroupFront
	groups := make([]structs.GroupFront, 0)
	tx, _ := Db.Begin()
	groupRows, err := tx.Query("select groupid,groupName,groupIntro,banner from `group`")
	if err != nil {
		fmt.Println("获取群组信息出错", err.Error())
	}

	for groupRows.Next() {
		err = groupRows.Scan(&group.ID, &group.Name, &group.Intro, &group.Banner)
		if err != nil {
			fmt.Println("读取SQL-写入群组数据出错", err.Error())
		}
		groups = append(groups, group)
	}
	for index, single := range groups {
		countrow := tx.QueryRow("select count(*) from post where groupid=?", single.ID)
		err = countrow.Scan(&groups[index].Amount)
		if err != nil {
			fmt.Println("读取SQL-写入群组count数据出错", err.Error())
		}
	}
	err = tx.Commit()
	if err != nil {
		fmt.Println("commit SQL出错", err.Error())
	}
	return groups
}

//NewGroup SQL 创建一个新的群组
func NewGroup(res structs.ResGroup) bool {
	tx, _ := Db.Begin()
	//如果是0，不存在此名称的群，可以插入
	ifrow := tx.QueryRow("select count(*) from `group` where groupName=?", res.Name)
	var ifexist int
	ifrow.Scan(&ifexist)
	if ifexist != 0 {
		return false
	}

	_, err := tx.Exec("insert into `group` (groupName,groupIntro,banner) values(?,?,?)",
		res.Name, res.Intro, res.Pic)
	if err != nil {
		fmt.Println("SQL插入群组错误", err.Error())
		return false
	}
	tx.Commit()
	return true

}
