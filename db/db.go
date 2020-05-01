package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var GormClient *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/kakao?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connection Established")
	}
	GormClient = db
}
