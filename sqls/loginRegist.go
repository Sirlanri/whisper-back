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
