package repository

import (
	"avito-shop/config"
	"avito-shop/models"
	"fmt"
	"gorm.io/gorm"
)

func CreateMerch(name string, price int) error {
	merch := models.Merch{Name: name, Price: price}
	return config.DB.Create(&merch).Error
}

func GetMerchByName(name string) (*models.Merch, error) {
	var merch *models.Merch
	err := config.DB.Where("name = ?", name).First(&merch).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return merch, fmt.Errorf("товар с именем '%s' не найден", name)
		}
		return merch, fmt.Errorf("ошибка при получении товара: %v", err)
	}

	return merch, nil
}

func BuyMerch(user *models.Merch, merch *models.Merch) error {
	return nil
}
