package persistence

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

func CreateDbConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	RunMigrations(db)
	conn = db

	return db
}

func GetConnection() *gorm.DB {
	return conn
}
