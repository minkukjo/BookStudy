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

func InsertUser(object model.User) {

	tx := GormClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	alreadyObj := object
	if FindFirstUser(alreadyObj) {
		tx.Model(&object).Debug().Update("token", object.Token)
		tx.Commit()
		return
	}

	if err := tx.Debug().Create(&object).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}

func FindFirstUser(object model.User) bool {

	tx := GormClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Debug().First(&object).RecordNotFound() {
		return false
	}

	tx.Rollback()

	return true
}
