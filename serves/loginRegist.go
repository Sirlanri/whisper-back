package serves

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"whisper/sqls"
)

//Login 传入mail和密码，验证密码是否正确，不正确返回false
func Login(mail string, pw string) (result bool, power int) {
	//hashed := Myhash(pw)
	pwFromDb, power := sqls.Login(mail)
	if Myhash(pw) == pwFromDb {
		fmt.Println("用户登录成功", mail)
		result = true
	} else {
		result = false
	}
	return
}

//Myhash 计算密码的哈希值
func Myhash(pw string) string {
	pw = pw + "IamRicoLan"
	myHash := md5.New()
	myHash.Write([]byte(pw))
	res := myHash.Sum(nil)
	result := hex.EncodeToString(res)
	return result
}

//Regist 接受用户名、邮箱、密码
func Regist(name, mail, pw string) (result string, code int) {
	if !Check(name, mail) {
		result = "用户名或邮箱格式不正确，请检查后输入"
		code = 202
	} else {
		result, code = sqls.Regist(name, mail, Myhash(pw))
	}
	return
}

//Check 正则表达式检验用户创建的用户名，邮箱是否合理（不与数据库查重） 合法返回true
func Check(name, mail string) bool {
	length := strings.Count(name, "")
	if length == 0 {
		return false
	}
	//检测邮箱
	mailres := regexp.MustCompile(`\S+@\S+\.`)
	result := mailres.MatchString(mail)
	if result {
		return true
	}
	return false

}
