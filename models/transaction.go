package models

type Transaction struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	FromUser string `gorm:"not null"`
	ToUser   string `gorm:"not null"`
	Amount   int    `gorm:"not null"`
}
