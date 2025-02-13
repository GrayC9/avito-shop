package repository

func BuyMerch(userName string, merchName string) error {
	//db := config.DB
	//
	//var user models.User
	//if err := db.Where("name = ?", userName).First(&user).Error; err != nil {
	//	return errors.New("пользователь не найден")
	//}
	//
	//var merch models.Merch
	//if err := db.Where("name = ?", merchName).First(&merch).Error; err != nil {
	//	return errors.New("товар не найден")
	//}
	//
	//if user.Coins < merch.Price {
	//	return errors.New("недостаточно монет")
	//}
	//
	//user.Coins -= merch.Price
	//if err := db.Save(&user).Error; err != nil {
	//	return errors.New("не удалось обновить баланс")
	//}
	//
	//transaction := models.Transaction{
	//	FromUser: user.Name,
	//	ToUser:   "shop",
	//	Amount:   merch.Price,
	//}
	//if err := db.Create(&transaction).Error; err != nil {
	//	return errors.New("ошибка сохранения транзакции")
	//}
	//
	//return nil
}
