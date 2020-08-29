package database

import "time"

type Users struct {
	Id           int64
	Email        string    `xorm:"notnull unique"`
	Password     string    `xorm:"notnull"`
	RegisteredAt time.Time `xorm:"notnull created"`
	Meta         string    `xorm:"text"`
	Role         int       `xorm:"tinyint notnull default 0"`
	DeletedAt    time.Time `xorm:"deleted"`
	Version      int       `xorm:"version"`
}
