package handlers

import (
	"avito-shop/models"
	"avito-shop/repository"
	"fmt"
	"net/http"
	"strings"
)

type BuyRequest struct {
	UserName  string `json:"user"`
	MerchName string `json:"merch"`
}

var merchant = &models.User{UserID: 6, Login: "shop"}

func BuyMerchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := UserAuthz(w, r)
	if err != nil {
		return
	}

	requestedItem := strings.TrimPrefix(r.URL.Path, "/api/buy/")
	merch, err := repository.GetMerchByName(requestedItem)
	if err != nil {
		http.Error(w, `{"error": "Некорректный тип товара"}`, http.StatusBadRequest)
		return
	}

	if user.Coins < merch.Price {
		http.Error(w, `{"error": "Недостаточно средств"}`, http.StatusPaymentRequired)
		return
	}

	err = repository.TransferCoins(user, merchant, merch.Price)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = repository.AccountMerchToUser(user, merch)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Покупка успешна"}`))
}
