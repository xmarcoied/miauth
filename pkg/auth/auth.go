package auth

import (
	"context"

	"github.com/xmarcoied/miauth/models"
	"github.com/xmarcoied/miauth/services/storage"
)

// AuthService defines the auth main service
type AuthService struct {
	storage storage.UsersInterface
}

func NewAuthService(userinterface storage.UsersInterface) *AuthService {
	return &AuthService{
		storage: userinterface,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, username, password string) (models.User, error) {
	hashPassword, err := s.GenerateHashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	return s.storage.CreateUser(ctx, username, hashPassword)
}

func (s *AuthService) Login()          {}
func (s *AuthService) UpdateUser()     {}
func (s *AuthService) ChangePassword() {}
func (s *AuthService) ResetPassword()  {}
