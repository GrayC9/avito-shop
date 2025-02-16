package migrations

import (
	"avito-shop/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func InitUserStatuses(db *gorm.DB) {
	statuses := []models.UserStatus{
		{UserStatusID: 1, Name: "active"},
		{UserStatusID: 2, Name: "blocked"},
		{UserStatusID: 3, Name: "fired"},
	}

	for _, status := range statuses {
		db.FirstOrCreate(&status, models.UserStatus{UserStatusID: status.UserStatusID})
	}
}

func InitShopUser(db *gorm.DB) {
	shopUser := models.User{
		Login:     "shop",
		Password:  "123",
		CreatedAt: time.Now(),
		StatusID:  2,
	}

	db.FirstOrCreate(&shopUser, models.User{Login: "shop"})
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Merch{},
		&models.Transaction{},
		&models.UserStatus{},
		&models.Token{},
		&models.MerchUser{},
		&models.Purchase{},
	)
	if err != nil {
		fmt.Println("❌ Ошибка миграции:", err)
	} else {
		fmt.Println("✅ База данных успешно мигрирована!")
	}
	InitUserStatuses(db)
	InitShopUser(db)
	InitMerch(db)
}

func InitMerch(db *gorm.DB) {
	merches := []models.Merch{
		{Name: "t-shirt", Price: 80},
		{Name: "cup", Price: 20},
		{Name: "book", Price: 50},
		{Name: "pen", Price: 10},
		{Name: "powerbank", Price: 200},
		{Name: "hoody", Price: 300},
		{Name: "umbrella", Price: 200},
		{Name: "socks", Price: 10},
		{Name: "wallet", Price: 50},
		{Name: "pink-hoody", Price: 500},
	}

	for _, merch := range merches {
		db.FirstOrCreate(&merch, models.Merch{Name: merch.Name})
	}

}
