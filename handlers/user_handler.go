package handlers

import (
	"net/http"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	//user, err := repository.GetUserByName("user1") // временно захардкоженный пользователь
	//if err != nil {
	//	http.Error(w, `{"errors": "Пользователь не найден"}`, http.StatusBadRequest)
	//	return
	//}
	//
	//response := map[string]interface{}{
	//	"coins": user.Coins,
	//}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(response)
}
