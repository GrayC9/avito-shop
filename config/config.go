package config

import (
	"avito-shop/migrations"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"db",
		"admin",
		"password",
		"avito_shop",
		"5432",
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ Не удалось подключиться к базе данных")
	}

	fmt.Println("✅ Успешное подключение к БД!")

	migrations.Migrate(DB)
}
