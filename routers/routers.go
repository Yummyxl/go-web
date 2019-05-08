package routers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"go-web/controller"
	"go-web/dto/response"
)

func InitRoute(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(response.BaseResponse(0, nil, "404 NOT FOUND"))
	})

	app.PartyFunc("/test", controller.TestRoute)

}
