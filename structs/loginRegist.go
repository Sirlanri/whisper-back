package structs

//ResLogin 从前端获取的登录json
type ResLogin struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

//ResRegist 从前端获取的注册json
type ResRegist struct {
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}
