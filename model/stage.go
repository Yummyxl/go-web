package model

import (
	"time"
)

type Stage struct {
	Id          int64     `xorm:"pk autoincr comment('主键id') BIGINT(20)"`
	IsDel       int       `xorm:"not null default 0 comment('是否删除') INT(11)"`
	Pid         int64     `xorm:"not null default 0 comment('学习阶段pid') BIGINT(20)"`
	Level       int       `xorm:"not null comment('级别 1,2') INT(11)"`
	Name        string    `xorm:"not null default '' comment('名称') VARCHAR(64)"`
	PreferIds   string    `xorm:"not null default '' comment('阶段兴趣ids串') VARCHAR(512)"`
	OtherStatus int64     `xorm:"not null default 1 comment('从最低位开始第一位是是否在boss展示，第二位是是否在c端 第三位是否是家长专区。第四位是否包含家长专区,第5位为是否展示兴趣标签') BIGINT(20)"`
	CreateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
}
