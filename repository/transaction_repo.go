package repository

import (
	"avito-shop/config"
	"avito-shop/models"
)

func AccountMerchToUser(user *models.User, merch *models.Merch) (*models.MerchUser, error) {
	merchUser := models.MerchUser{
		UserID:  user.UserID,
		MerchID: merch.MerchID,
	}
	return &merchUser, config.DB.Create(&merchUser).Error
}
