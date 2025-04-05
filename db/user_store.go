package db

import (
	"context"

	"github.com/kave08/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName  = "hotel-reservation"
	userCol = "users"
)

type UserStore interface {
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	GetUsers(ctx context.Context) ([]types.User, error)
	CreateUser(ctx context.Context, user *types.User) (*types.User, error)
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
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	if err := m.Col.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *MongoUserStore) GetUsers(ctx context.Context) ([]types.User, error) {
	var users []types.User
	cur, err := m.Col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.Decode(&users); err != nil {
		return []types.User{}, nil
	}

	return users, nil
}

func (m *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	return nil, nil
}
