package model

type User struct {
	Id          int    `gorm:"primary_key" json:"id"`
	ConnectedAt string `gorm:"size:255" json:"connected_at"`
	Nickname    string `gorm:"size:255" json:"nickname"`
	Token       string `gorm:"size:255" json:"token"`
}
