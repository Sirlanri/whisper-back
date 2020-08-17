package serves

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"whisper/sqls"
)

//Login 传入mail和密码，验证密码是否正确，不正确返回false
func Login(mail string, pw string) (result bool, power int) {
	//hashed := Myhash(pw)
	pwFromDb, power := sqls.Login(mail)
	if pw == pwFromDb {
		fmt.Println("用户登录成功", mail)
		result = true
	} else {
		result = false
	}
	return
}

//Myhash 计算密码的哈希值
func Myhash(pw string) string {
	afterHash := md5.New().Sum([]byte(pw))
	after64 := base64.StdEncoding.EncodeToString(afterHash)
	return after64
}

//Regist 接受用户名、邮箱、密码
func Regist(name, mail, pw string) (result string) {

}
