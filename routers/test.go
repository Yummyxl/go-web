package routers

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

func TestRouter(party router.Party)  {
	party.Get("/ping", func(context context.Context) {
		context.WriteString("pong")
	})
}
