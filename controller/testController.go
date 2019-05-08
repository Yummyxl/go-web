package controller

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

func TestRoute(test router.Party) {

	test.Get("/ping", func(context context.Context) {
		context.WriteString("pong")
	})
}
