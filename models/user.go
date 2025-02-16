package models

import "time"

type User struct {
	UserID    uint      `gorm:"column:user_id;primaryKey;autoIncrement"`
	Login     string    `gorm:"column:login;type:text;not null"`
	Coins     int       `gorm:"column:coins;type:int;not null;default:1000"`
	Password  string    `gorm:"column:password;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:date;not null"`
	StatusID  uint      `gorm:"column:status_id"`
}

func (User) TableName() string {
	return "users"
}

type UserStatus struct {
	UserStatusID int    `gorm:"column:user_status_id;primaryKey;autoIncrement"`
	Name         string `gorm:"column:name;type:text;not null"`
}

func (UserStatus) TableName() string {
	return "user_statuses"
}

type Token struct {
	TokenID   int       `gorm:"column:token_id;primaryKey;autoIncrement"`
	UserID    int       `gorm:"column:user_id;not null"`
	Token     string    `gorm:"column:token;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	ExpiredAt time.Time `gorm:"column:expired_at;type:timestamp"`
}

func (Token) TableName() string {
	return "tokens"
}

type Purchase struct {
	PurchaseID uint      `gorm:"column:purchase_id;primaryKey;autoIncrement"`
	UserID     uint      `gorm:"column:user_id;not null"`
	MerchID    uint      `gorm:"column:merch_id;not null"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (Purchase) TableName() string {
	return "purchases"
}
