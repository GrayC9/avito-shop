package repository

import (
	"avito-shop/config"
	"avito-shop/models"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AccountMerchToUser(user *models.User, merch *models.Merch) error {
	purchase := models.Purchase{
		UserID:  user.UserID,
		MerchID: merch.MerchID,
	}
	if err := config.DB.Create(&purchase).Error; err != nil {
		return errors.New("failed to record purchase")
	}
	return nil
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
