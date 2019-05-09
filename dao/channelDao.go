package dao

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"go-web/database"
	"go-web/model"
)

type ChannelDao struct {
	engine *xorm.Engine
}

func NewChannelDao() *ChannelDao {
	return &ChannelDao{
		engine: database.NewDatabase(),
	}
}

func (dao *ChannelDao) SelectChannelById(id int64) *model.Channel {
	channel := new(model.Channel)
	channel.Id = id
	bo, e := dao.engine.Get(channel)
	fmt.Println(*channel)
	if !bo || e != nil {
		return nil
	} else {
		return channel
	}
}
