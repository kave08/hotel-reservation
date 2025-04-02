package db

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	dbName = "hotel-reservation"
)

func ToObjectID(id string) primitive.ObjectID {
	bid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return bid
}
