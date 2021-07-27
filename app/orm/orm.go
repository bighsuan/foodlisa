package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase(dsn string) (db *gorm.DB) {

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("使用 GORM 連線資料庫發生錯誤, 原因為 " + err.Error())
	}

	// TODO: iterate all struct in models and init table

	return
}


