package repository

//
//import (
//	"avito-shop/config"
//	"avito-shop/models"
//)
//
//func CreateUser(name string) error {
//	user := models.User{Name: name, Coins: 1000}
//	return config.DB.Create(&user).Error
//}
//
//func GetUserByName(name string) (models.User, error) {
//	var user models.User
//	err := config.DB.Where("name = ?", name).First(&user).Error
//	return user, err
//}
//
//func UpdateUserCoins(name string, newCoins int) error {
//	return config.DB.Model(&models.User{}).Where("name = ?", name).Update("coins", newCoins).Error
//}
