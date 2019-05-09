package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go-web/controller"
	"go-web/dto/response"
)

func InitRoute(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(response.BaseResponse(iris.StatusNotFound, nil, "404 NOT FOUND"))
	})
	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(response.BaseResponse(iris.StatusInternalServerError, nil, "ERROR"))
	})
	app.PartyFunc("/test", controller.TestRoute)
	app.PartyFunc("/channel", controller.ChannelRoute)
}
