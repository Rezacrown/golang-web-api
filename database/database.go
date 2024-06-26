package database

import (
	"go-web-api/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInstance() *gorm.DB {

	// database url
	dsn := "root:@tcp(127.0.0.1:3306)/golang_web_api?charset=utf8mb4&parseTime=True&loc=Local"

	// membuka koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migrasi table
	db.AutoMigrate(&book.Book{})

	// println("database sukses konek")

	return db
}
