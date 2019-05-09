package model

import (
	"time"
)

type StageExt struct {
	Id          int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	IsDel       int       `xorm:"not null default 0 comment('是否删除') INT(11)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	StageId     int64     `xorm:"not null default 0 comment('阶段id') BIGINT(20)"`
	EntityType  int       `xorm:"not null default 0 comment('类型 0频道 1 广告') INT(11)"`
	EntityValue int64     `xorm:"not null comment('值') BIGINT(20)"`
}
