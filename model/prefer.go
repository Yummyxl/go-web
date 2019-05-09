package model

import (
	"time"
)

type Prefer struct {
	Id         int64     `xorm:"pk autoincr comment('主键') BIGINT(20)"`
	Name       string    `xorm:"not null default '' comment('名称') VARCHAR(32)"`
	CreateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}
