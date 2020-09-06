package structs

//ResPost 发布新post，前端发送的json
type ResPost struct {
	Content string   `json:"content"`
	Pics    string   `json:"pics"`
	Group   string   `json:"group"`
	Tags    []string `json:"tags"`
}
