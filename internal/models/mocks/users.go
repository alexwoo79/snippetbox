package mocks

import (
	"time"

	"github.com/alexwoo79/snippetbox/internal/models"
)

func (m *models.UserModel) Get(id int) (models.User, error) {
	if id == 1 {
		u := models.User{
			ID:      1,
			Name:    "Alex",
			Email:   "alex@example.com",
			Created: time.Now(),
		}
		return u, nil
	}
	return models.User{}, models.ErrNoRecord
}

func (m *models.UserModel) PasswordUpdate(id int, currentPassword, newPassword string) error {
	if id == 1 {
		if currentPassword != "pa$$word" {
			return models.ErrInvalidCredentials
		}
		return nil
	}
	return models.ErrNoRecord
}
