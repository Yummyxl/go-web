package service

import (
	"go-web/dao"
	"go-web/dto/response/channel"
)

type ChannelService interface {
	GetChannelById(id int64) *channel.ChannelGetByIdResponse
}

type channelService struct {
	dao *dao.ChannelDao
}

func NewChannelService() ChannelService {
	return &channelService{
		dao: dao.NewChannelDao(),
	}
}

func (service *channelService) GetChannelById(id int64) (response *channel.ChannelGetByIdResponse) {
	res := service.dao.SelectChannelById(id)
	response = &channel.ChannelGetByIdResponse{
		Id:          res.Id,
		DisplayName: res.DisplayName,
		NoteName:    res.NoteName,
		IsHandpick:  res.IsHandpick,
	}
	return
}
