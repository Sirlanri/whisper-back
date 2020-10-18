package structs

//ReplyDetail 从数据库中获取的完整reply信息
type ReplyDetail struct {
	ID       int    `json:"id"` //被回复的reply的id
	Postid   int    `json:"postid"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	HaveRead bool   `json:"haveRead"`
	Content  string `json:"content"`
}

//ResNewReply 前端传入的新回复
type ResNewReply struct {
	ID      int    `json:"id"`      //被回复post的id
	Content string `json:"content"` //回复的内容
	Name    string `json:"name"`    //被回复人的name
}
