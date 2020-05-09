package model

type Post struct {
	Id     int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY" json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Date   string `json:"date"`
	Text   string `gorm:"type:text" json:"text"`
	Name   string `json:"name"`
	Kind   string `json:"kind"`
}
