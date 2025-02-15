package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword_Success(t *testing.T) {
	password := "securepassword"
	hashedPassword, err := HashPassword(password)

	assert.NoError(t, err)

	assert.NotEqual(t, password, hashedPassword)
}

func TestCheckPasswordHash_Success(t *testing.T) {
	password := "securepassword"
	hashedPassword, _ := HashPassword(password)

	assert.True(t, CheckPasswordHash(password, hashedPassword))
}

func TestCheckPasswordHash_Failure(t *testing.T) {
	password := "securepassword"
	incorrectPassword := "wrongpassword"
	hashedPassword, _ := HashPassword(password)

	assert.False(t, CheckPasswordHash(incorrectPassword, hashedPassword))
}
