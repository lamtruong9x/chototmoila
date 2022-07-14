package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	USER   = "root"
	PASS   = "pass"
	HOST   = "localhost"
	PORT   = "3030"
	DBNAME = "cho_tot"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicln("Cannot close DB")
	}
	// Close
	sqlDB.Close()
}
