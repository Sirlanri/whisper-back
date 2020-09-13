package structs

//ResPost 发布新post，前端发送的json
type ResPost struct {
	Content string   `json:"content"`
	Pics    []string `json:"pics"`
	Group   string   `json:"group"`
	Tags    []string `json:"tags"`
}

//DataPost 从数据库中获取的完整post
type DataPost struct {
	ID      int      `json:"id"`
	Avatar  string   `json:"avatar"`
	User    string   `json:"username"`
	Group   string   `json:"groupname"`
	GroupID int      `json:"groupid"`
	Content string   `json:"content"`
	Topic   []string `json:"topic"`
	Time    string   `json:"time"`
	Pics    []string `json:"pics"`
	Replys  []Reply  `json:"replys"`
}

//Reply 完整post中的回复结构体
type Reply struct {
	Name    string `json:"name"`
	Imgsrc  string `json:"imgsrc"`
	Content string `json:"content"`
}
