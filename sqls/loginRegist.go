package sqls

//Login SQL 从数据库中获取邮箱对应的密码，返回给serve
func Login(mail string) string {
	var pw string
	row := Db.QueryRow("select password from user where mail=?", mail)
	row.Scan(&pw)
	return pw
}
