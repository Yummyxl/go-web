package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main()  {

	app := iris.Default()

	app.Handle("GET", "/ping", func(context context.Context) {
		context.WriteString("pong")
	})

	app.Run(iris.Addr(":8080"))
}
