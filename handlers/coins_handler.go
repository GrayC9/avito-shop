package handlers

import (
	"avito-shop/models"
	"avito-shop/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

type SendCoinRequest struct {
	ToUser string `json:"toUser" binding:"required"`
	Amount int    `json:"amount" binding:"required"`
}

func UserAuthz(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	authz := r.Header.Get("Authorization")
	userID, err := TokenIdentify(authz)
	if err != nil && userID == -1 {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("Unauthorized"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return nil, err
		}

		w.WriteHeader(http.StatusUnauthorized)
		return nil, err
	}
	user, err := repository.GetUserById(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil, err
	}
	//if user.UserID == userID {
	//	w.WriteHeader(http.StatusUnauthorized)
	//	return nil, err
	//}

	return &user, nil
}

func SendCoinsHandler(w http.ResponseWriter, r *http.Request) {
	// Проверить запрос
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Авторизовать пользователя
	userFrom, err := UserAuthz(w, r)
	// Нужно что-то в логи писать, если не авторизован юзер? По идее нет, потому что это базовое поведение
	// Было бы хорошо сюда воткнуть метрику прометея для дальнейшего мониторинга, но пока оставим тело условия пустым.
	if err != nil {
		return
	}

	// Спарсить тело запроса и опредеелить получателя
	var req SendCoinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Некорректное тело запроса"}`, http.StatusBadRequest)
		return
	}

	userTo, err := repository.GetUserByLogin(req.ToUser)
	if err != nil {
		http.Error(w, `{"error": "Получатель не найден"}`, http.StatusNotFound)
		return
	}

	// Отправить монеты
	err = repository.TransferCoins(userFrom, userTo, req.Amount)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()), http.StatusPaymentRequired)
		return
	}

	// Совершить транзакцию
	w.Header().Set("Content-Type", "application/json")
	_, err = fmt.Fprintf(w, "Coins moved")
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Получить стоимость товара
// Проверить баланс пользователя
// Совершить транзакцию
// Записать товар на баланс пользователя
// Вернуть ок
