package mongo

import (
	"context"
	"errors"

	"github.com/xmarcoied/miauth/models"
	"github.com/xmarcoied/miauth/services/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	if s.IsUserExist(ctx, user.Username) {
		return models.User{}, storage.ErrAlreadyExist
	}
	s.client.Database("testing").Collection("users").
		InsertOne(ctx, user)

	return models.User{}, nil
}
func (s *Service) GetUser(ctx context.Context, username string) (models.User, error) {
	user := models.User{}
	err := s.client.Database("testing").Collection("users").
		FindOne(ctx, bson.D{{"username", username}}).Decode(&user)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.User{}, storage.ErrNotFound
		}
		return models.User{}, err
	}
	return user, nil
}
func (s *Service) UpdateUser(ctx context.Context, username string, user models.User) error {
	_, err := s.client.Database("testing").Collection("users").
		UpdateOne(ctx, bson.D{{"username", username}}, bson.M{"$set": user})

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) IsUserExist(ctx context.Context, username string) bool {
	err := s.client.Database("testing").Collection("users").
		FindOne(ctx, bson.D{{"username", username}}).Err()

	return !errors.Is(err, mongo.ErrNoDocuments)
}
