package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(hash2("cococooc" + "thisIsRico!"))
	fmt.Println(hash2("123456" + "thisIsRico!"))
	fmt.Println(hash2("123456"))

}

//Myhash 计算密码的哈希值
func Myhash(pw string) string {
	afterHash := md5.New().Sum([]byte(pw))
	after64 := base64.StdEncoding.EncodeToString(afterHash)
	return after64
}

func hash2(pw string) string {
	myHash := md5.New()
	myHash.Write([]byte(pw))
	res := myHash.Sum(nil)
	result := hex.EncodeToString(res)
	return result
}
