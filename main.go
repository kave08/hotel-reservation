package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/kave08/hotel-reservation/api"
	"github.com/kave08/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbURI   = "mongodb://localhost:27017"
	dbName  = "hotel-reservation"
	userCol = "users"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	user := types.User{
		FirstName: "kave",
		LastName:  "hudj",
	}

	coll := client.Database(dbName).Collection(userCol)
	_, err = coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return
	}

	var kave types.User

	err = coll.FindOne(ctx, bson.M{}).Decode(&kave)
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	err = app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}
