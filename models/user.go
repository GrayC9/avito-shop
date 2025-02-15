package models

import "time"

type User struct {
	UserID    int       `gorm:"column:user_id;primaryKey;autoIncrement"`
	Login     string    `gorm:"column:login;type:text;not null"`
	Coins     int       `gorm:"column:coins;type:int;not null"`
	Password  string    `gorm:"column:password;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:date;not null"`
	StatusID  int       `gorm:"column:status_id"`
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
	UserID    int       `gorm:"column:user_id;not null"` // Ссылается на пользователя
	Token     string    `gorm:"column:token;type:text;not null"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	ExpiredAt time.Time `gorm:"column:expired_at;type:timestamp"`
}

func (Token) TableName() string {
	return "tokens"
}
