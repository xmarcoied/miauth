package auth

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xmarcoied/miauth/models"
	"github.com/xmarcoied/miauth/pkg"
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
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	hashPassword, err := s.GenerateHashPassword(password)
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": username,
		}).Error("cannot create a new user")
		return models.User{}, err
	}
	user, err := s.storage.CreateUser(ctx, username, hashPassword)
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": username,
		}).Error("cannot create a new user")
		return models.User{}, err
	}

	pkg.GetLogContext(ctx).WithFields(log.Fields{
		"user": username,
	}).Info("new user created")

	return user, nil
}

// Login connects a user against datastore
func (s *Service) Login() error {
	return nil
}

// UpdateUser updates user info
func (s *Service) UpdateUser() {}

// ChangePassword changes user's password
func (s *Service) ChangePassword() {}

// ResetPassword resets user's password to a new randomized password
func (s *Service) ResetPassword() {}
