package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return
}
