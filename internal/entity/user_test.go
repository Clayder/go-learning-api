package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	
	t.Run("should create a new user", func(t *testing.T) {
		user, err := NewUser("Clayder", "test@gmail.com", "123456")
		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "Clayder", user.Name)
		assert.Equal(t, "test@gmail.com", user.Email)
		assert.NotEmpty(t, user.Password)
	})

	t.Run("should validate password", func(t *testing.T) {
		password := "123456"
		user, err := NewUser("Clayder", "test@gmail.com", password)
		assert.Nil(t, err)
		assert.True(t, user.ValidatePassword(password))
		assert.False(t, user.ValidatePassword("wrongpassword"))
		assert.NotEqual(t, password, user.Password)
	})
}