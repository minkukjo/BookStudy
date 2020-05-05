package model

type Post struct {
	Id     int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Text   string `json:"text"`
	Name   string `json:"name"`
}
