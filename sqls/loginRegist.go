package sqls

//Login SQL 从数据库中获取邮箱对应的密码，返回给serve
func Login(mail string) (string, int) {
	var (
		pw    string
		power int
	)
	row := Db.QueryRow("select password,power from user where mail=?", mail)
	row.Scan(&pw, &power)
	return pw, power
}

//Regist SQL 先查询是否存在重复的邮箱或昵称，再插入
