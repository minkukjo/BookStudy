package db

import (
	"log"
	"testing"
)

type User struct {
	Id          int    `gorm:"primary_key;auto_increment" json:"id"`
	ConnectedAt string `gorm:"size:255" json:"connected_at"`
	Nickname    string `gorm:"size:255;not null;unique" json:"nickname"`
}

func TestGorm(t *testing.T) {
	user := &User{
		ConnectedAt: "2020-04-24",
		Nickname:    "Harry",
	}

	user2 := &User{
		ConnectedAt: "2020-04-24",
		Nickname:    "minkuk",
	}

	tx := GormClient.Begin()
	defer tx.Rollback()

	err := tx.Debug().AutoMigrate(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Debug().Create(&user).Error
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Debug().Create(&user2).Error
	if err != nil {
		log.Fatal(err)
	}

	err = tx.Debug().First(&user).Error
	if err != nil {
		log.Fatal(err)
	}
}
