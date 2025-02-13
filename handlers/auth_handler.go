package handlers

import (
	"encoding/json"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"errors": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{"token": "mocked_jwt_token"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
