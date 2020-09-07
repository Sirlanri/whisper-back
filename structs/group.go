package structs

//GroupFront 前端接收的群组json
type GroupFront struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Name   string `json:"name"`
	Intro  string `json:"intro"`
	Bannar string `json:"jmgsrc"`
}
