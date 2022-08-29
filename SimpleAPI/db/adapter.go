package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	dbDSN = "root:password@tcp(gorm_db)/gorm_todo?charset=utf8mb4&parseTime=True&loc=Local"
)

func Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", os.Getenv("GORM_DIALECT"))
	return db, err
}