package handlers

import (
	"avito-shop/models"
	"avito-shop/repository"
	"avito-shop/utils"
	"encoding/json"
	"net/http"
)

type InfoResponse struct {
	Coins       int                         `json:"coins"`
	Inventory   []models.Merch              `json:"inventory"`
	CoinHistory map[string][]models.History `json:"coinHistory"` // Используем models.History
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := utils.ExtractUserIDFromToken(r)
	if err != nil {
		http.Error(w, `{"error": "Unauthorized"}`, http.StatusUnauthorized)
		return
	}

	user, err := repository.GetUserById(userID)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		return
	}

	received, sent, err := repository.GetTransactionHistory(userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch transactions"}`, http.StatusInternalServerError)
		return
	}

	inventory, err := repository.GetUserInventory(userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to fetch inventory"}`, http.StatusInternalServerError)
		return
	}

	response := InfoResponse{
		Coins:     user.Coins,
		Inventory: inventory,
		CoinHistory: map[string][]models.History{
			"received": received,
			"sent":     sent,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
