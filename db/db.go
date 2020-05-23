package db

import (
	"bookstudy/config"
	"bookstudy/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var GormClient *gorm.DB

func init() {
	var db *gorm.DB
	var err error
	if *(config.Config) == "local" {
		db, err = gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/kakao?charset=utf8&parseTime=True")
	} else {
		db, err = gorm.Open("mysql", "root:1234@tcp(mysqldb)/kakao?charset=utf8&parseTime=True")
	}

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connection Established")
	}
	GormClient = db

	if err := GormClient.Debug().AutoMigrate(&model.User{}, &model.Post{}).Error; err != nil {
		log.Fatal(err)
	}
}

func CreateUser(object model.User) {

	tx := GormClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	alreadyObj := object
	if FindFirstUser(&alreadyObj) {
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

func FindFirstUser(object *model.User) bool {

	tx := GormClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Debug().First(object).RecordNotFound() {
		return false
	}

	tx.Rollback()

	return true
}

func CreatePost(object *model.Post) {
	tx := GormClient.Begin()

	if err := tx.Debug().Create(&object).Error; err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
}

func FindAllPosts(query string, arg string) []model.Post {

	var posts []model.Post
	tx := GormClient.Begin()

	if err := tx.Debug().Where(query+" = ?", arg).Find(&posts).Error; err != nil {
		tx.Rollback()
		return nil
	}

	tx.Commit()

	return posts
}

func DeletePost(id int) error {
	post := model.Post{
		Id: id,
	}

	tx := GormClient.Begin()

	if err := tx.Debug().Delete(&post).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func UpdatePost(object *model.Post) error {

	tx := GormClient.Begin()

	if err := tx.Model(&object).Debug().
		Update(model.Post{
			Title: object.Title,
			Text:  object.Text,
			Kind:  object.Kind,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil

}
