package infra

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/takuya911/mcrsvc_user/conf"
)

// NewGormDB function
func NewGormDB() (*gorm.DB, error) {
	mysqlConf := &mysql.Config{
		User:                 conf.C.DB.User,
		Passwd:               conf.C.DB.Pass,
		Net:                  "tcp",
		Addr:                 conf.C.DB.Host + ":" + conf.C.DB.Port,
		DBName:               conf.C.DB.Name,
		ParseTime:            true,
		Collation:            conf.C.DB.Collation,
		Loc:                  time.Local,
		AllowNativePasswords: true,
	}

	db, err := gorm.Open("mysql", mysqlConf.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
