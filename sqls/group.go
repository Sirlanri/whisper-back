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
	groupRows, err := tx.Query("select groupid,groupName,groupIntro,banner from igroup")
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
	ifrow := tx.QueryRow("select count(*) from igroup where groupName=?", res.Name)
	var ifexist int
	ifrow.Scan(&ifexist)
	if ifexist != 0 {
		return false
	}

	_, err := tx.Exec("insert into igroup (groupName,groupIntro,banner) values(?,?,?)",
		res.Name, res.Intro, res.Pic)
	if err != nil {
		fmt.Println("SQL插入群组错误", err.Error())
		return false
	}
	tx.Commit()
	return true

}

//NewGroup2 SQL 创建一个新的群组 用一个SQL完成 速度竟然不如上面的！
func NewGroup2(res structs.ResGroup) bool {
	tx, _ := Db.Begin()
	pre, err := Db.Prepare("INSERT INTO igroup" + ` (groupName,groupIntro,banner)
		SELECT ?,?,?
		from DUAL
		WHERE not exists (
				SELECT * ` +
		"from  igroup" +
		"WHERE groupName=?);")
	if err != nil {
		fmt.Println("预编译表达式出错", err.Error())
	}
	_, err = pre.Exec(res.Name, res.Intro, res.Pic, res.Name)
	if err != nil {
		fmt.Println("SQL插入群组错误", err.Error())
		return false
	}
	tx.Commit()
	return true

}

/*DelGroupOnly SQL
删除一个群，保留并修改其post*/
func DelGroupOnly(groupid int) bool {
	tx, _ := Db.Begin()

	//删除群
	_, err := tx.Exec(`delete from igroup where groupid=?`, groupid)
	if err != nil {
		fmt.Println("SQL 删除群出错", err.Error())
	}

	//将群内的post全部修改状态
	numsRow, err := tx.Exec(`update post set groupid=0 where groupid=?`, groupid)
	if err != nil {
		fmt.Println("修改群post的id出错", err.Error())
		return false
	}
	nums, _ := numsRow.RowsAffected()
	fmt.Printf("删除的群ID：%d，影响了%d条post", groupid, nums)

	tx.Commit()
	return true
}
