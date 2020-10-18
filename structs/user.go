package structs

//UserInfo 获取用户信息返回的数据类型
type UserInfo struct {
	Name       string `json:"name"`
	Mail       string `json:"mail"`
	Intro      string `json:"intro"`
	Avatar     string `json:"avatar"`
	Bannar     string `json:"bannar"`
	PostCount  int    `json:"postCount"`
	ReplyCount int    `json:"replyCount"`
	Power      string `json:"power"`
}

/*ResChangeInfo JSON
修改资料，前端传入的信息*/
type ResChangeInfo struct {
	Name  string `json:"name"`
	Mail  string `json:"mail"`
	Intro string `json:"intro"`
}
