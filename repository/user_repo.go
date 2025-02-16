package repository

import (
	"avito-shop/config"
	"avito-shop/models"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetUserInventory(userID int) ([]models.Merch, error) {
	var inventory []models.Merch

	err := config.DB.Model(&models.Merch{}).
		Joins("LEFT JOIN purchases p ON p.merch_id = merch.merch_id").
		Where("user_id = ?", userID).Find(&inventory).Error

	return inventory, err
}

func GetUserByLogin(login string) (*models.User, error) {
	var user *models.User
	err := config.DB.Where("login = ? and status_id = 1", login).First(&user).Error
	return user, err
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	err := config.DB.Where("user_id = ? and status_id = 1", id).First(&user).Error
	return user, err
}

func GetUserBalance(user *models.User) (int, error) {
	var u models.User
	result := config.DB.Model(&models.User{}).Select("coins").Where("login = ?", user.Login).First(&u)
	if result.Error != nil {
		return -1, result.Error
	}

	return u.Coins, nil
}

func TransferCoins(userFrom *models.User, userTo *models.User, amount int) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		balance, err := GetUserBalance(userFrom)
		if err != nil {
			return err
		}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&userFrom, userFrom.UserID).Error; err != nil {
			return err
		}

		if balance < amount {
			return errors.New("Недостаточно монет")
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&userTo, userTo.UserID).Error; err != nil {
			return err
		}

		if err := tx.Model(&userFrom).
			Update("coins", userFrom.Coins-amount).Error; err != nil {
			return err
		}

		if err := tx.Model(&userTo).
			Update("coins", userTo.Coins+amount).Error; err != nil {
			return err
		}

		tr := &models.Transaction{
			FromUserID: int(userFrom.UserID),
			ToUserID:   int(userTo.UserID),
			Amount:     amount,
		}

		if err := tx.Model(tr).Save(tr).Error; err != nil {
			return err
		}

		return nil
	})

}
