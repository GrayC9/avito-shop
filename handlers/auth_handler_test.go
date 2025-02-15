package handlers

import (
	"avito-shop/config"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Тестируем успешную генерацию токена
func TestTokenHandler_Success(t *testing.T) {
	// Мокаем подключение к базе данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Устанавливаем mock для функции DB
	dsn := "host=localhost user=admin password=password dbname=avito_shop port=5432 sslmode=disable" // Пример DSN
	config.DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Создаем пользователя
	// Ожидаем, что запрос к базе данных будет успешным
	mock.ExpectQuery("SELECT (.+) FROM \"users\" WHERE \"login\" = ?").
		WithArgs("test_user").
		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "password", "created_at", "status_id"}).
			AddRow(1, "test_user2", "hashed_password", time.Now(), 1))

	// Ожидаем создание токена
	mock.ExpectExec("INSERT INTO \"tokens\"").
		WithArgs(1, "generated_token", time.Now(), time.Now().Add(24*time.Hour)).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Подготавливаем запрос
	reqBody := `{"login": "test_user2", "password": "password2"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Создаем запись ответа
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	TokenHandler(rr, req)

	// Проверяем статус-код
	assert.Equal(t, http.StatusOK, rr.Code)

	// Проверяем, что ответ содержит токен
	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Contains(t, response, "token")
}

// Тестируем ошибку, если пароль не верный
func TestTokenHandler_InvalidPassword(t *testing.T) {
	// Мокаем подключение к базе данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка при создании мока базы данных: %v", err)
	}
	defer db.Close()

	// Устанавливаем mock для функции DB
	dsn := "host=localhost user=admin password=password dbname=avito_shop port=5432 sslmode=disable" // Пример DSN
	config.DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Создаем пользователя
	// Ожидаем, что запрос к базе данных будет успешным
	mock.ExpectQuery("SELECT (.+) FROM \"users\" WHERE \"login\" = ?").
		WithArgs("test_user").
		WillReturnRows(sqlmock.NewRows([]string{"id", "login", "password", "created_at", "status_id"}).
			AddRow(1, "test_user", "hashed_password", time.Now(), 1))

	// Подготавливаем запрос с неверным паролем
	reqBody := `{"login": "test_user", "password": "wrong_password"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth", bytes.NewBufferString(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Создаем запись ответа
	rr := httptest.NewRecorder()

	// Вызываем обработчик
	TokenHandler(rr, req)

	// Проверяем статус-код
	assert.Equal(t, http.StatusUnauthorized, rr.Code)

	// Проверяем сообщение об ошибке
	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, response["error"], "Неверный логин или пароль")
}
