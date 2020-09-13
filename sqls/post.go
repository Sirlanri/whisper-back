package sqls

import (
	"fmt"
	"time"
	"whisper/structs"
)

//NewPost 负责处理前端接收到的数据
func NewPost(res structs.ResPost, userid int) {
	var (
		err     error
		groupid int
	)
	//计时
	t1 := time.Now()
	tx, _ := Db.Begin()

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

//GetALlPost SQL 获取全部post
func GetALlPost() (posts []structs.DataPost) {
	tx, _ := Db.Begin()
	var (
		userids []int
		userid  int
		post    structs.DataPost
		reply   structs.Reply //replys列表的单个回复元素
	)

	//SQL获取全部post
	postsRow, err := tx.Query(`SELECT * FROM post ORDER BY postid DESC LIMIT 20`)
	if err != nil {
		fmt.Println("查询post列表出错", err.Error())
	}

	//Scan回复
	for postsRow.Next() {
		err = postsRow.Scan(&post.ID, &userid, &post.GroupID, &post.Content, &post.Time)
		if err != nil {
			fmt.Println("SQL 读取后写入post出错", err.Error())
		}
		post.Time = post.Time[5:16]
		userids = append(userids, userid)
		posts = append(posts, post)
	}

	for index, singlePost := range posts {
		//根据userid获取用户昵称+头像，写入posts
		userRow := tx.QueryRow(`select userName,avatar from user where userid=?`,
			userids[index])
		err = userRow.Scan(&posts[index].User, &posts[index].Avatar)
		if err != nil {
			fmt.Println("SQL 写入user信息出错", err.Error())
		}

		//通过groupid获取群名称
		if posts[index].GroupID == 0 {
			posts[index].Group = ""
		} else {
			groupNameRow := tx.QueryRow(`select groupName from igroup where groupid=?`, posts[index].GroupID)
			groupNameRow.Scan(&posts[index].Group)
		}

		//通过postid获取topic
		topicRow, err := tx.Query(`select topic from tag where postid=?`, singlePost.ID)
		if err != nil {
			fmt.Println("获取topic失败", err.Error())
		} else {
			for topicRow.Next() {
				var topic string
				topicRow.Scan(&topic)
				posts[index].Topic = append(posts[index].Topic, topic)
			}
		}

		//获取图片
		var picrows []string
		picRow, _ := tx.Query(`select picaddress from picture where postid=?`, singlePost.ID)
		for picRow.Next() {
			var pic string
			picRow.Scan(&pic)
			picrows = append(picrows, pic)
		}
		posts[index].Pics = picrows

		//获取replys
		var (
			replyUserids []int
			replyUserid  int
			replys       []structs.Reply //单个post的回复列表
		)
		replysRow, err := tx.Query(`select fromUser,content
		 from reply where postid=?`, singlePost.ID)
		if err != nil {
			fmt.Println("SQL 通过postID获取replys出错", err.Error())
		}
		for replysRow.Next() {
			err = replysRow.Scan(&replyUserid, &reply.Content)
			if err != nil {
				fmt.Println("SQL 读取后写入reply出错", err.Error())
			}
			replyUserids = append(replyUserids, replyUserid)
			replys = append(replys, reply)
		}

		//通过replyUserid获取用户昵称+头像
		for index, userid := range replyUserids {
			userRows := tx.QueryRow(`select userName,avatar from user where userid=?`,
				userid)
			if err != nil {
				fmt.Println("SQL 通过id读取user信息出错", err.Error())
			}
			//单个reply信息已完善，添加至replys列表
			userRows.Scan(&replys[index].Name, &replys[index].Imgsrc)
		}

		//将整理好的replys添加至post
		posts[index].Replys = replys
	}
	tx.Commit()
	return

}

//GetALlPostByUser SQL 获取某个用户的全部post
func GetALlPostByUser(name string) (posts []structs.DataPost) {
	tx, _ := Db.Begin()
	var (
		userids []int
		userid  int
		post    structs.DataPost
		reply   structs.Reply //replys列表的单个回复元素
	)
	var gotUserid int
	//通过name获取用户id
	idRow := tx.QueryRow(`select userid from user where userName=?`, name)
	idRow.Scan(&gotUserid)
	//SQL获取全部post
	postsRow, err := tx.Query(`SELECT * FROM post where publisher=? ORDER BY postid DESC LIMIT 20`, gotUserid)
	if err != nil {
		fmt.Println("查询post列表出错", err.Error())
	}

	//Scan回复
	for postsRow.Next() {
		err = postsRow.Scan(&post.ID, &userid, &post.GroupID, &post.Content, &post.Time)
		if err != nil {
			fmt.Println("SQL 读取后写入post出错", err.Error())
		}
		post.Time = post.Time[5:16]
		userids = append(userids, userid)
		posts = append(posts, post)
	}

	for index, singlePost := range posts {
		//根据userid获取用户昵称+头像，写入posts
		userRow := tx.QueryRow(`select userName,avatar from user where userid=?`,
			userids[index])
		err = userRow.Scan(&posts[index].User, &posts[index].Avatar)
		if err != nil {
			fmt.Println("SQL 写入user信息出错", err.Error())
		}

		//通过groupid获取群名称
		if posts[index].GroupID == 0 {
			posts[index].Group = ""
		} else {
			groupNameRow := tx.QueryRow(`select groupName from igroup where groupid=?`, posts[index].GroupID)
			groupNameRow.Scan(&posts[index].Group)
		}

		//通过postid获取topic
		topicRow, err := tx.Query(`select topic from tag where postid=?`, singlePost.ID)
		if err != nil {
			fmt.Println("获取topic失败", err.Error())
		} else {
			for topicRow.Next() {
				var topic string
				topicRow.Scan(&topic)
				posts[index].Topic = append(posts[index].Topic, topic)
			}
		}

		//获取图片
		var picrows []string
		picRow, _ := tx.Query(`select picaddress from picture where postid=?`, singlePost.ID)
		for picRow.Next() {
			var pic string
			picRow.Scan(&pic)
			picrows = append(picrows, pic)
		}
		posts[index].Pics = picrows

		//获取replys
		var (
			replyUserids []int
			replyUserid  int
			replys       []structs.Reply //单个post的回复列表
		)
		replysRow, err := tx.Query(`select fromUser,content
		 from reply where postid=?`, singlePost.ID)
		if err != nil {
			fmt.Println("SQL 通过postID获取replys出错", err.Error())
		}
		for replysRow.Next() {
			err = replysRow.Scan(&replyUserid, &reply.Content)
			if err != nil {
				fmt.Println("SQL 读取后写入reply出错", err.Error())
			}
			replyUserids = append(replyUserids, replyUserid)
			replys = append(replys, reply)
		}

		//通过replyUserid获取用户昵称+头像
		for index, userid := range replyUserids {
			userRows := tx.QueryRow(`select userName,avatar from user where userid=?`,
				userid)
			if err != nil {
				fmt.Println("SQL 通过id读取user信息出错", err.Error())
			}
			//单个reply信息已完善，添加至replys列表
			userRows.Scan(&replys[index].Name, &replys[index].Imgsrc)
		}

		//将整理好的replys添加至post
		posts[index].Replys = replys
	}
	tx.Commit()
	return
}

/*GetPostByGroup SQL
传入群组id，返回该群组的post 限制20*/
func GetPostByGroup(groupid int) (posts []structs.DataPost) {
	tx, _ := Db.Begin()

	var (
		userids []int
		userid  int
		post    structs.DataPost
		reply   structs.Reply //replys列表的单个回复元素
	)
	postsRow, err := tx.Query(`SELECT * FROM post where groupid=?
	 ORDER BY postid DESC LIMIT 25`, groupid)
	if err != nil {
		fmt.Println("通过groupid查询post出错")
	}
	//Scan回复
	for postsRow.Next() {
		err = postsRow.Scan(&post.ID, &userid, &post.GroupID, &post.Content, &post.Time)
		if err != nil {
			fmt.Println("SQL 读取后写入post出错", err.Error())
		}
		post.Time = post.Time[5:16]
		userids = append(userids, userid)
		posts = append(posts, post)
	}

	for index, singlePost := range posts {
		//根据userid获取用户昵称+头像，写入posts
		userRow := tx.QueryRow(`select userName,avatar from user where userid=?`,
			userids[index])
		err = userRow.Scan(&posts[index].User, &posts[index].Avatar)
		if err != nil {
			fmt.Println("SQL 写入user信息出错", err.Error())
		}

		//通过groupid获取群名称
		if posts[index].GroupID == 0 {
			posts[index].Group = ""
		} else {
			groupNameRow := tx.QueryRow(`select groupName from igroup where groupid=?`, posts[index].GroupID)
			groupNameRow.Scan(&posts[index].Group)
		}

		//通过postid获取topic
		topicRow, err := tx.Query(`select topic from tag where postid=?`, singlePost.ID)
		if err != nil {
			fmt.Println("获取topic失败", err.Error())
		} else {
			for topicRow.Next() {
				var topic string
				topicRow.Scan(&topic)
				posts[index].Topic = append(posts[index].Topic, topic)
			}
		}

		//获取图片
		var picrows []string
		picRow, _ := tx.Query(`select picaddress from picture where postid=?`, singlePost.ID)
		for picRow.Next() {
			var pic string
			picRow.Scan(&pic)
			picrows = append(picrows, pic)
		}
		posts[index].Pics = picrows

		//获取replys
		var (
			replyUserids []int
			replyUserid  int
			replys       []structs.Reply //单个post的回复列表
		)
		replysRow, err := tx.Query(`select fromUser,content
		 from reply where postid=?`, singlePost.ID)
		if err != nil {
			fmt.Println("SQL 通过postID获取replys出错", err.Error())
		}
		for replysRow.Next() {
			err = replysRow.Scan(&replyUserid, &reply.Content)
			if err != nil {
				fmt.Println("SQL 读取后写入reply出错", err.Error())
			}
			replyUserids = append(replyUserids, replyUserid)
			replys = append(replys, reply)
		}

		//通过replyUserid获取用户昵称+头像
		for index, userid := range replyUserids {
			userRows := tx.QueryRow(`select userName,avatar from user where userid=?`,
				userid)
			if err != nil {
				fmt.Println("SQL 通过id读取user信息出错", err.Error())
			}
			//单个reply信息已完善，添加至replys列表
			userRows.Scan(&replys[index].Name, &replys[index].Imgsrc)
		}

		//将整理好的replys添加至post
		posts[index].Replys = replys
	}
	tx.Commit()
	return
}
