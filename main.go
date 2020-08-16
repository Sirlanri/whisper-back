package main

import (
	"whisper/handlers"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, handlers.NotFound)
}
