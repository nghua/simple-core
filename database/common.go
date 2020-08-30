package database

import "time"

type Users struct {
	Id           int64     `xorm:"notnull pk"`
	Email        string    `xorm:"notnull unique"`
	Password     string    `xorm:"notnull"`
	RegisteredAt time.Time `xorm:"notnull created"`
	Meta         string    `xorm:"text notnull default ''"`
	Role         int       `xorm:"tinyint notnull default 0"`
	DeletedAt    time.Time `xorm:"deleted"`
	Version      int       `xorm:"version"`
}

type Terms struct {
	Id        int64
	Name      string    `xorm:"notnull"`
	Count     int64     `xorm:"notnull default 0"`
	Meta      string    `xorm:"text notnull default ''"`
	Type      int       `xorm:"tinyint notnull default 0"`
	DeletedAt time.Time `xorm:"deleted"`
	Version   int       `xorm:"version"`
}

type TermRelationships struct {
	Id   int64
	Term int64 `xorm:"notnull"`
	Post int64 `xorm:"notnull"`
}
