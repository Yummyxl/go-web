package routers

import (
	"github.com/kataras/iris/core/router"
	"go-web/controller"
)

func Routers(routers router.Party) {

	routers.PartyFunc("/test", controller.TestRoute)

}
