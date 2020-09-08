package sqls

import (
	"fmt"
	"time"
	"whisper/structs"

	"google.golang.org/protobuf/internal/errors"
)

//NewPost 负责处理前端接收到的数据
func NewPost(res structs.ResPost, mail string) {
	var (
		err             error
		userid, groupid int
	)
	//计时
	t1 := time.Now()

	tx, _ := Db.Begin()
	//通过邮箱获取用户id
	idRow := tx.QueryRow("select userid from user where mail=?", mail)
	err = idRow.Scan(&userid)
	if err != nil {
		fmt.Println("SQL获取用户ID出错", err.Error())
	}
	fmt.Println("获取用户ID ", time.Since(t1))

	//获取群组id
	groupRow := tx.QueryRow("select groupid from igroup where groupName=?", res.Group)
	err = groupRow.Scan(&groupid)
	if err != nil {
		fmt.Println("获取群ID出错", err.Error())
	}
	fmt.Println("获取群组ID ", time.Since(t1))

	//写入post表
	pre, err := tx.Prepare(`INSERT INTO post (publisher,groupid,content)
	 values (?,?,?)`)
	if err != nil {
		fmt.Println("预编译SQL出错", err.Error())
	}
	newrow, err := pre.Exec(userid, groupid, res.Content)
	if err != nil {
		fmt.Println("SQL写入post出错 ", err.Error())
	}
	fmt.Println("写入post ", time.Since(t1))
	postid, _ := newrow.LastInsertId()

	//插入picture表
	for _, pic := range res.Pics {
		_, err = tx.Exec("insert into picture values (?,?)", postid, pic)
		if err != nil {
			fmt.Println("SQL 插入图片出错", err.Error())
		}
	}
	fmt.Println("写入图片 ", time.Since(t1))

	//写入tag
	for _, tag := range res.Tags {
		_, err = tx.Exec("insert into tag values (?,?)", tag, postid)
		if err != nil {
			fmt.Println("SQL 插入Tag出错", err.Error())
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("SQL commit出错 ", err.Error())
	}
	elapsed := time.Since(t1)
	fmt.Println("共计耗时 ", elapsed)
	fmt.Println("SQL写入完毕")
}

//NewPost2 另一种读取数据库的方式，更慢了mmp
func NewPost2(res structs.ResPost, mail string) {
	var (
		err             error
		userid, groupid int
	)
	//计时
	t1 := time.Now()

	//通过邮箱获取用户id
	idRow := Db.QueryRow("select userid from user where mail=?", mail)
	err = idRow.Scan(&userid)
	if err != nil {
		fmt.Println("SQL获取用户ID出错", err.Error())
	}
	fmt.Println("获取用户ID ", time.Since(t1))

	//获取群组id
	groupRow := Db.QueryRow("select groupid from igroup where groupName=?", res.Group)
	err = groupRow.Scan(&groupid)
	if err != nil {
		fmt.Println("获取群ID出错", err.Error())
	}
	fmt.Println("获取群组ID ", time.Since(t1))

	//写入post表
	pre, err := Db.Prepare(`INSERT INTO post (publisher,groupid,content)
	 values (?,?,?)`)
	if err != nil {
		fmt.Println("预编译SQL出错", err.Error())
	}
	newrow, err := pre.Exec(userid, groupid, res.Content)
	if err != nil {
		fmt.Println("SQL写入post出错 ", err.Error())
	}
	fmt.Println("写入post ", time.Since(t1))
	postid, _ := newrow.LastInsertId()

	//插入picture表
	for _, pic := range res.Pics {
		_, err = Db.Exec("insert into picture values (?,?)", postid, pic)
		if err != nil {
			fmt.Println("SQL 插入图片出错", err.Error())
		}
	}
	fmt.Println("写入图片 ", time.Since(t1))

	if err != nil {
		fmt.Println("SQL commit出错 ", err.Error())
	}
	elapsed := time.Since(t1)
	fmt.Println("共计耗时 ", elapsed)
	fmt.Println("SQL写入完毕")
}

//GetGroupNames SQL 从数据库获取全部的群组名称 用于创建post
func GetGroupNames() (groups []string) {
	tx, _ := Db.Begin()
	groupRows, err := tx.Query("select groupName from igroup")
	if err != nil {
		fmt.Println("SQL 获取Group列表出错", err.Error())
	}
	var groupName string
	for groupRows.Next() {
		err = groupRows.Scan(&groupName)
		if err != nil {
			fmt.Println("SQL scan出错", err.Error())
		}
		groups = append(groups, groupName)
	}
	fmt.Println(groups)
	return
}

//GetTags SQL 从数据库获取全部的tag
func GetTags() (tags []string) {
	tx, _ := Db.Begin()
	tagRows, err := tx.Query("select distinct topic from `tag`")
	if err != nil {
		fmt.Println("SQL 获取Tag列表出错", err.Error())
	}
	var tagName string
	for tagRows.Next() {
		err = tagRows.Scan(&tagName)
		if err != nil {
			fmt.Println("SQL scan出错", err.Error())
		}
		tags = append(tags, tagName)
	}
	fmt.Println(tags)
	return
}

//GetAllPost SQL 获取全部的post
func GetAllPost() (posts []structs.DataPost) {
	tx, _ := Db.Begin()
	//这边限制查询数量，以后会懒加载
	postRow, err := tx.Query("SELECT * FROM post ORDER BY postid DESC LIMIT 20")
	if err != nil {
		fmt.Println("SQL 查找posts失败", errors.Error())
	}
	//单个回复post
	var post structs.DataPost
	for postRow.Next() {
		err = postRow.Scan(&post.ID, &post.User, &post.Group, &post.Content, &post.Time)
		if err != nil {
			fmt.Println("SQL 读取后写入post出错", err.Error())
		}
		posts = append(posts, post)
	}
	//评论列表,以便写入posts
	var replylist []structs.Reply
	for index, single := range posts {
		//获取tag
		tagRow, err := tx.Query("select topic from tag where postid=?", single.ID)
		if err != nil {
			fmt.Println("SQL 获取tag失败/无tag", err.Error())
		}
		var topicList []string
		for tagRow.Next() {
			var oneTopic string
			tagRow.Scan(&oneTopic)
			topicList = append(topicList, oneTopic)
		}
		posts[index].Topic = topicList

		//获取评论
		var (
			single   structs.Reply
			singleid int
		)
		replyRow, err := tx.Query(`select fromUser,content
		from reply where postid=?`, single.ID)
		//临时存放获取的用户id列表
		var idlist []int
		for replyRow.Next() {
			replyRow.Scan(&singleid, &posts[index].Content)
			idlist = append(idlist, singleid)
		}
		//获取评论中的name

		for index, id := range idlist {
			nameRow := tx.QueryRow("select userName from user where userid=?", id)
			err = nameRow.Scan(&single.Name)
			if err != nil {
				fmt.Println("Scan nameRow出错", err.Error())
			}

		}
	}
}
