package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"go-web/routers"
)

func main()  {

	app := iris.New()
	app.Use(recover2.New())
	app.Use(logger.New())

	app.PartyFunc("/test", routers.TestRouter)
	app.Run(iris.Addr(":8080"))
}
