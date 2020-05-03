package db

import (
	"bookstudy/model"
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

	if err := GormClient.Debug().AutoMigrate(&model.User{}).Error; err != nil {
		log.Fatal(err)
	}
}

func Insert(object interface{}) {

	tx := GormClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Debug().Create(object).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
		return
	}

	tx.Commit()
}

func FindFirst(object interface{}) {

	tx := GormClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Debug().First(object).RecordNotFound() {
		log.Println("못찾음")
	}

	if err := tx.Debug().First(object).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
		return
	}

	tx.Commit()
}
