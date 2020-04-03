package app

import (
	"github.com/jinzhu/gorm"
	"log"
)

func openDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Create(item interface{}) {
	db := openDB()
	defer db.Close()

	db.AutoMigrate(&item)

	db.Create(&item)
}

func Read() {

}
