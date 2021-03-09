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

func NewService(userinterface storage.UsersInterface) *Service {
	return &Service{
		storage: userinterface,
	}
}

func (s *Service) CreateUser(ctx context.Context, username, password string) (models.User, error) {
	hashPassword, err := s.GenerateHashPassword(password)
	if err != nil {
		return models.User{}, err
	}
	return s.storage.CreateUser(ctx, username, hashPassword)
}

func (s *Service) Login()          {}
func (s *Service) UpdateUser()     {}
func (s *Service) ChangePassword() {}
func (s *Service) ResetPassword()  {}
