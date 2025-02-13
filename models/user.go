package models

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null"`
	Coins int    `gorm:"default:1000"`
}
