package mongo

import (
	"context"

	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Service struct {
	client *mongo.Client
}

func New(ctx context.Context) *Service {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	return &Service{
		client: client,
	}
}

func (s *Service) Shutdown() {
	log.Warn("shutting down mongo connection")
	if err := s.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
