package auth

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/xmarcoied/miauth/models"
	"github.com/xmarcoied/miauth/pkg"
	"github.com/xmarcoied/miauth/pkg/rand"
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
func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	hashPassword, err := s.GenerateHashPassword(req.Password)
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": req.Username,
		}).Error("cannot create a new user")
		return models.User{}, err
	}

	user, err := s.storage.CreateUser(ctx, models.User{
		Username:  req.Username,
		Password:  hashPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": req.Username,
		}).Error("cannot create a new user")
		return models.User{}, err
	}

	pkg.GetLogContext(ctx).WithFields(log.Fields{
		"user": req.Username,
	}).Info("new user created")

	return user, nil
}

// Login connects a user against datastore
func (s *Service) Login(ctx context.Context, req LoginRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	user, err := s.storage.GetUser(ctx, req.Username)
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": req.Username,
		}).Error("user is not found")
		return err
	}

	isValid, _ := s.IsHashPasswordValid(user.Password, req.Password)
	if !isValid {
		return errors.New("password is not valid")
	}

	return nil
}

// UpdateUser updates user info
func (s *Service) UpdateUser(ctx context.Context, username string, req UpdateUserRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	user, err := s.storage.GetUser(ctx, username)
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": username,
		}).Error("user is not found")
		return err
	}

	log.Println(req)

	err = s.storage.UpdateUser(ctx, username, models.User{
		Username:  user.Username,
		Password:  user.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})

	return err
}

// ChangePassword changes user's password
func (s *Service) ChangePassword() error {
	return nil
}

// ResetPassword resets user's password to a new randomized password
func (s *Service) ResetPassword(ctx context.Context, username string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	user, err := s.storage.GetUser(ctx, username)
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": username,
		}).Error("user is not found")
		return "", err
	}

	newpassword := rand.String(10)
	hashPassword, err := s.GenerateHashPassword(newpassword)
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": username,
		}).Error("cannot create a new user")
		return "", err
	}

	err = s.storage.UpdateUser(ctx, username, models.User{
		Username:  username,
		Password:  hashPassword,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		pkg.GetLogContext(ctx).WithError(err).WithFields(log.Fields{
			"user": username,
		}).Error("user cannot update")
		return "", err
	}

	return newpassword, nil
}
