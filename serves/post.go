package serves

import uuid "github.com/satori/go.uuid"

//Createid 为图片生成唯一名称
func Createid() string {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4(), nil)
	id := u1.String()
	return id[:9]
}
