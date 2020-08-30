package database

import (
	"fmt"
	"simple-core/public/setting"

	log "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
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

	err = engine.Sync2(new(Users), new(Terms), new(TermRelationships))
	if err != nil {
		log.Fatalf("同步数据库出错：%v", err)
	}

	_, err = engine.Exec(`CREATE TABLE IF NOT EXISTS uuid (
    h int(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    x tinyint(4) NOT NULL UNIQUE DEFAULT '0'
)`)
	if err != nil {
		log.Fatalf("同步数据库出错：%v", err)
	}

	engine.ShowSQL(true)
}

func Engine() *xorm.Engine {
	return engine
}
