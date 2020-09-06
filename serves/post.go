package serves

import (
	"whisper/structs"

	uuid "github.com/satori/go.uuid"
)

//Createid 为图片生成唯一名称
func Createid() string {
	// 创建 UUID v4
	u1 := uuid.Must(uuid.NewV4(), nil)
	id := u1.String()
	return id[:9]
}

//NewPost 负责处理前端接收到的数据
func NewPost(ResPost structs.ResPost) {

}
