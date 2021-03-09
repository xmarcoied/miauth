package storage

import (
	"context"

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

func (m *MockedUser) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	if m.IsUserExist(user.Username) {
		return models.User{}, ErrAlreadyExist
	}

	m.datastore = append(m.datastore, user)

	return user, nil
}

func (m *MockedUser) GetUser(ctx context.Context, username string) (models.User, error) {
	for _, user := range m.datastore {
		if user.Username == username {
			return user, nil
		}
	}

	return models.User{}, ErrNotFound
}

func (m *MockedUser) UpdateUser(ctx context.Context, username string, user models.User) error {
	for i, user := range m.datastore {
		if user.Username == username {
			m.datastore[i].FirstName = user.FirstName
			m.datastore[i].LastName = user.LastName
			return nil
		}
	}

	return ErrNotFound
}

func (m *MockedUser) IsUserExist(username string) bool {
	for _, user := range m.datastore {
		if user.Username == username {
			return true
		}
	}

	return false
}
