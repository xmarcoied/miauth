package auth

import (
	"context"

	"github.com/xmarcoied/miauth/models"
	"github.com/xmarcoied/miauth/services/storage"
)

// Service defines the auth main service
type Service struct {
	storage storage.UsersInterface
}

// NewService initiatize the main Auth service with storage interface
func NewService(userinterface storage.UsersInterface) *Service {
	return &Service{
		storage: userinterface,
	}
}

// CreateUser create a new user and return the created user info
func (s *Service) CreateUser(ctx context.Context, username, password string) (models.User, error) {
	hashPassword, err := s.GenerateHashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	return s.storage.CreateUser(ctx, username, hashPassword)
}

// Login connects a user against datastore
func (s *Service) Login() {}

// UpdateUser updates user info
func (s *Service) UpdateUser() {}

// ChangePassword changes user's password
func (s *Service) ChangePassword() {}

// ResetPassword resets user's password to a new randomized password
func (s *Service) ResetPassword() {}
