package handlers

import (
	"avito-shop/repository"
	"encoding/json"
	"net/http"
)

type BuyRequest struct {
	UserName  string `json:"user"`
	MerchName string `json:"merch"`
}

func BuyMerchHandler(w http.ResponseWriter, r *http.Request) {
	var req BuyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Неверный формат запроса"}`, http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByName(req.UserName)
	if err != nil {
		http.Error(w, `{"error": "Пользователь не найден"}`, http.StatusBadRequest)
		return
	}

	merch, err := repository.GetMerchByName(req.MerchName)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	if user.Coins < merch.Price {
		http.Error(w, `{"error": "Недостаточно монет для покупки"}`, http.StatusBadRequest)
		return
	}

	err = repository.BuyMerch(req.UserName, req.MerchName)
	if err != nil {
		http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	newBalance := user.Coins - merch.Price
	err = repository.UpdateUserCoins(req.UserName, newBalance)
	if err != nil {
		http.Error(w, `{"error": "Не удалось обновить баланс"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Покупка успешна"}`))
}
