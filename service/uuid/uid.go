package uuid

import (
	"database/sql"
	"simple-core/database"

	"github.com/edwingeng/wuid/mysql/wuid"
)

var uid *wuid.WUID

func getDb() (*sql.DB, bool, error) {
	db, err := database.Engine().NewDB()
	return db.DB, true, err
}

func init() {
	uid = wuid.NewWUID("simple-uid", nil)
}

func GenerateUid() (int64, error) {
	wuid.WithSection(6)
	err := uid.LoadH28FromMysql(getDb, "uuid")
	if err != nil {
		return 0, err
	}
	return uid.Next(), nil
}
