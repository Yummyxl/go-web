package model

import (
	"time"
)

type Channel struct {
	Id          int64     `xorm:"pk autoincr comment('主键id') BIGINT(11)"`
	IsDel       int       `xorm:"not null default 0 comment('是否删除') INT(11)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	DisplayName string    `xorm:"not null default '' comment('外显名称') VARCHAR(16)"`
	NoteName    string    `xorm:"not null default '' comment('备注名称') VARCHAR(20)"`
	IsHandpick  int       `xorm:"not null default 0 comment('是否精选') INT(11)"`
}
