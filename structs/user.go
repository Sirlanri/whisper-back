package structs

//UserInfo 获取用户信息返回的数据类型
type UserInfo struct {
	Name       string `json:"name"`
	Intro      string `json:"intro"`
	Avatar     string `json:"avatar"`
	Bannar     string `json:"bannar"`
	PostCount  int    `json:"postcount"`
	ReplyCount int    `json:"replycount"`
}
