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

type TokenRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Некорректный JSON"}`, http.StatusBadRequest)
		return
	}

	var user models.User
	result := config.DB.Where("login = ?", req.Login).First(&user)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		http.Error(w, `{"error": "Ошибка базы данных"}`, http.StatusInternalServerError)
		return
	}

	if result.RowsAffected > 0 {
		http.Error(w, `{"error": "Пользователь уже существует"}`, http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		http.Error(w, `{"error": "Ошибка хеширования пароля"}`, http.StatusInternalServerError)
		return
	}

	user = models.User{
		Login:     req.Login,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		StatusID:  1,
	}
	config.DB.Create(&user)

	token, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		http.Error(w, `{"error": "Ошибка генерации токена"}`, http.StatusInternalServerError)
		return
	}

	tokenRecord := models.Token{
		UserID:    user.UserID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}
	config.DB.Create(&tokenRecord)

	resp := RegisterResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error": "Метод не поддерживается"}`, http.StatusMethodNotAllowed)
		return
	}

	var req TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error": "Некорректный JSON"}`, http.StatusBadRequest)
		return
	}

	var user models.User
	result := config.DB.Where("login = ?", req.Login).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, `{"error": "Неверный логин или пароль"}`, http.StatusUnauthorized)
			return
		}
		http.Error(w, `{"error": "Ошибка базы данных"}`, http.StatusInternalServerError)
		return
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		http.Error(w, `{"error": "Неверный логин или пароль"}`, http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.UserID)
	if err != nil {
		http.Error(w, `{"error": "Ошибка генерации токена"}`, http.StatusInternalServerError)
		return
	}

	tokenRecord := models.Token{
		UserID:    user.UserID,
		Token:     token,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}
	config.DB.Create(&tokenRecord)

	resp := TokenResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
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
