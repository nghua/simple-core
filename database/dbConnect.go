package database

import (
	"fmt"
	"simple-core/setting"

	log "github.com/sirupsen/logrus"

	"xorm.io/xorm"
)

var engine *xorm.Engine

func init() {
	var err error
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",
		setting.DbUser, setting.DbPassword, setting.DbHost, setting.DbName)
	engine, err = xorm.NewEngine("mysql", source)
	if err != nil {
		log.Fatalf("数据库连接出错：%v", err)
	}

	err = engine.Sync2(new(Users))
	if err != nil {
		log.Fatalf("同步数据库出错：%v", err)
	}

	engine.ShowSQL(true)
}

func Engine() *xorm.Engine {
	return engine
}
