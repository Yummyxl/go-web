package model

import (
	"time"
)

type ChannelExt struct {
	Id          int64     `xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	IsDel       int       `xorm:"not null default 0 comment('是否删除') INT(11)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	ChannelId   int64     `xorm:"not null comment('频道id') BIGINT(20)"`
	EntityType  int       `xorm:"not null comment('类型0 广告 1 公开课 2 付费课 3 类目 ') INT(11)"`
	EntityValue int64     `xorm:"not null comment('对应type的值') BIGINT(20)"`
}
