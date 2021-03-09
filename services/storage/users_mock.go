package storage

import (
	"context"
	"errors"

	"github.com/xmarcoied/miauth/models"
)

// MockedUser mocked the UsersInterface
type MockedUser struct {
	datastore []models.User
}

func NewMockedUser(store []models.User) *MockedUser {
	return &MockedUser{
		datastore: store,
	}
}

func (m *MockedUser) CreateUser(ctx context.Context, username, password string) (models.User, error) {
	if m.IsUserExist(username) {
		return models.User{}, errors.New("Already exist")
	}

	user := models.User{
		Username: username,
		Password: "password",
	}
	m.datastore = append(m.datastore, user)

	return user, nil
}

func (m *MockedUser) IsUserExist(username string) bool {
	for _, user := range m.datastore {
		if user.Username == username {
			return true
		}
	}

	return false
}
