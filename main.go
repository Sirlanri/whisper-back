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

	app.Run(iris.Addr(":8090"))

	return
}
