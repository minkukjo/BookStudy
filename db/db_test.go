package db

import (
	"fmt"
	"testing"
)

type User struct {
	Id          int    `gorm:"primary_key;auto_increment" json:"id"`
	ConnectedAt string `gorm:"size:255" json:"connected_at"`
	Nickname    string `gorm:"size:255;" json:"nickname"`
	Token       string `gorm:"size:255;" json:"token"`
}

var user = &User{
	Id:          1234,
	ConnectedAt: "2020-04-24",
	Nickname:    "Harry",
	Token:       "asdfsdfasfsf",
}

// JUnit처럼 Transaction으로 테스트하는 방법이 없으려나
func TestInsert(t *testing.T) {

	//user2 := &User{
	//	Id:          5678,
	//	ConnectedAt: "2020-04-24",
	//	Nickname:    "minkuk",
	//	Token:       "ggawtwerqrwqr",
	//}

	tx := GormClient.Begin()

	tx.Debug().Create(&user)
}

func TestFindFirst(t *testing.T) {
	tx := GormClient.Begin()

	tx.Debug().Create(&user)

	var target User

	tx.Debug().First(&target)
	fmt.Println(target)
}
