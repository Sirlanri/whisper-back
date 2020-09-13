package main

import (
	"whisper/handlers"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, handlers.NotFound)
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	whisper := app.Party("/whisper", crs).AllowMethods(iris.MethodOptions)

	whisper.Post("/login", handlers.Login)
	whisper.Post("/regist", handlers.Regist)
	whisper.Get("/logout", handlers.Logout)

	whisper.Get("/getUserInfoByName", handlers.GetUserInfoByName)
	whisper.Get("/getUserInfoByCookie", handlers.GetUserInfoByCookie)
	whisper.Get("/getGroupNames", handlers.GetGroupNames)
	whisper.Get("/getTags", handlers.GetTags)
	whisper.Get("/getGroups", handlers.GetGroups)
	whisper.Get("/getAllPost", handlers.GetAllPost)
	whisper.Get("/getPostByUser", handlers.GetPostByUser)
	whisper.Get("/getPostByGroup", handlers.GetPostByGroup)
	whisper.Get("/getAllReply", handlers.GetAllReply)

	whisper.Get("/readMsg", handlers.ReadMsg)

	whisper.Get("/changeAvatar", handlers.ChangeAvatar)
	whisper.Get("/changeBannar", handlers.ChangeBannar)
	whisper.Post("/changeInfo", handlers.ChangeInfo)

	whisper.Post("/uploadPics", handlers.UploadPics)
	whisper.Post("/newPost", handlers.NewPost)
	whisper.Post("/newGroup", handlers.NewGroup)
	whisper.Post("/newReply", handlers.NewReply)

	whisper.HandleDir("/getpics", iris.Dir("./uploadpics"))
	app.Run(iris.Addr(":8090"))

	return
}
