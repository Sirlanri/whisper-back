package structs

//GroupFront 前端接收的群组json
type GroupFront struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Name   string `json:"name"`
	Intro  string `json:"intro"`
	Banner string `json:"imgsrc"`
}

//ResGroup 创建新群组，从前端接收的数据
type ResGroup struct {
	Name  string `json:"name"`
	Intro string `json:"intro"`
	Pic   string `json:"pic"`
}
