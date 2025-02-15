package models

type Transaction struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	FromUserID int  `gorm:"not null"`
	ToUserID   int  `gorm:"not null"`
	Amount     int  `gorm:"not null"`
}

func (Transaction) TableName() string {
	return "transactions"
}
