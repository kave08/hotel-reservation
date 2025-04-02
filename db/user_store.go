package db

import (
	"context"

	"github.com/kave08/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCol = "users"
)

type UserStore interface {
	GetUserByID(ctx context.Context, id string) (*types.User, error)
}

type MongoUserStore struct {
	Client *mongo.Client
	Col    *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		Client: client,
		Col:    client.Database(dbName).Collection(userCol),
	}
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := m.Col.FindOne(ctx, bson.M{"_id": ToObjectID(id)}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
