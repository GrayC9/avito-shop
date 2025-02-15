package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	password := "securepassword"
	hashedPassword, err := HashPassword(password)

	// Проверяем, что хеширование прошло без ошибок
	assert.NoError(t, err)

	// Проверяем, что хешированный пароль не равен исходному
	assert.NotEqual(t, password, hashedPassword)
}

func TestCheckPasswordHash_Success(t *testing.T) {
	password := "securepassword"
	hashedPassword, _ := HashPassword(password)

	// Проверяем, что пароль совпадает с хешом
	assert.True(t, CheckPasswordHash(password, hashedPassword))
}

func TestCheckPasswordHash_Failure(t *testing.T) {
	password := "securepassword"
	incorrectPassword := "wrongpassword"
	hashedPassword, _ := HashPassword(password)

	// Проверяем, что неверный пароль не совпадает с хешом
	assert.False(t, CheckPasswordHash(incorrectPassword, hashedPassword))
}
