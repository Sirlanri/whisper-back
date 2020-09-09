package structs

//ReplyDetail 从数据库中获取的完整reply信息
type ReplyDetail struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	HaveRead bool   `json:"haveRead"`
	Content  string `json:"content"`
}
