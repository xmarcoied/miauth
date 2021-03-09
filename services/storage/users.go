package storage

import (
	"context"

	"github.com/xmarcoied/miauth/models"
)

// UsersInterface abstracts the User database storage to allow multiple implementations
type UsersInterface interface {
	CreateUser(ctx context.Context, username, password string) (models.User, error)
	GetUser(ctx context.Context, username string) (models.User, error)
}
