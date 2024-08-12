package model

type Post struct {
	Id          int    `gorm:"primaryKey, autoIncrement" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm: "not null" json:"description"`
}
