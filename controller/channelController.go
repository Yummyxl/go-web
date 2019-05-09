package controller

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	channel2 "go-web/dto/request/channel"
	"go-web/dto/response"
	"go-web/service"
)

var (
	channelService service.ChannelService
)

func ChannelRoute(party router.Party) {

	party.Post("/getChannelById", func(c context.Context) {
		channel := new(channel2.ChannelGetByIdRequest)
		c.ReadJSON(channel)
		if channelService == nil {
			channelService = service.NewChannelService()
		}

		c.JSON(response.Success(channelService.GetChannelById(channel.Id)))
	})
}
