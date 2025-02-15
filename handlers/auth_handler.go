package handlers

import (
	"avito-shop/config"
	"avito-shop/models"
	"avito-shop/utils"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type AuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	var req AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Некорректный JSON"})
		return
	}

	if req.Login == "" || req.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Логин и пароль обязательны"})
		return
	}

	var user models.User
	result := config.DB.Where("login = ?", req.Login).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			hashedPassword, err := utils.HashPassword(req.Password)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(ErrorResponse{Error: "Ошибка хеширования пароля"})
				return
			}

			user = models.User{
				Login:     req.Login,
				Password:  hashedPassword,
				CreatedAt: time.Now(),
				StatusID:  1,
			}
			config.DB.Create(&user)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Ошибка базы данных"})
			return
		}
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Неверный логин или пароль"})
		return
	}

	token, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Ошибка генерации токена"})
		return
	}

	tokenRecord := models.Token{
		UserID:    user.UserID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}
	config.DB.Create(&tokenRecord)

	resp := AuthResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func RevokeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, `{"error": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Токен не предоставлен"})
		return
	}

	claims, err := utils.ParseJWT(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Неверный токен"})
		return
	}

	var tokenRecord models.Token
	result := config.DB.Where("token = ? AND expired_at > ?", tokenString, time.Now()).First(&tokenRecord)
	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Токен не найден или просрочен"})
		return
	}

	if tokenRecord.UserID != claims.UserID {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Токен не принадлежит этому пользователю"})
		return
	}

	var user models.User
	result = config.DB.Where("user_id = ? AND status_id = 1", claims.UserID).First(&user)
	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Пользователь не найден или заблокирован"})
		return
	}

	resp := map[string]string{"message": "Токен действителен и пользователь активен"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func TokenIdentify(t string) (int, error) {
	var token models.Token
	result := config.DB.Select("user_id").Where("token = ? and expired_at > ?", t, time.Now()).First(&token)
	if result.Error != nil {
		return -1, result.Error
	}

	return token.UserID, nil
}
