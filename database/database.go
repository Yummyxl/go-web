package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var (
	Engine *xorm.Engine
)

func NewDatabase() *xorm.Engine {
	if Engine != nil {
		return Engine
	}
	engine, err := xorm.NewEngine("mysql", "chaikuservice:testchaiku@(172.21.139.13:3306)/haoke?charset=utf8")
	if err != nil {
		log.Fatal("创建database失败", err)
	}

	engine.ShowSQL(true)
	Engine = engine
	return Engine
}
