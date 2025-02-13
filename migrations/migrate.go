package migrations

import (
	"avito-shop/models"
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Merch{}, &models.Transaction{})
	if err != nil {
		fmt.Println("❌ Ошибка миграции:", err)
	} else {
		fmt.Println("✅ База данных успешно мигрирована!")
	}
}
