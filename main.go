package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"go-web/dto/response"
	"go-web/routers"
)

func main() {
	app := iris.New()
	app.Use(recover2.New())
	app.Use(logger.New())

	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(response.BaseResponse(0, nil, "404 NOT FOUND"))
	})

	app.PartyFunc("/", routers.Routers)

	app.Run(iris.Addr(":8080"))
}
