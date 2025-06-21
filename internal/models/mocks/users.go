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
