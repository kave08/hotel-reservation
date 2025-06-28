package main

import (
	"context"
	"flag"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/kave08/hotel-reservation/db"
	"github.com/kave08/hotel-reservation/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var config = fiber.Config{
	ErrorHandler: func(c fiber.Ctx, err error) error {
		return c.JSON(map[string]string{
			"error": err.Error(),
		})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	dbName := os.Getenv("MONGO_DATABASE")
	dbURI := "mongodb://" + username + ":" + password + "@localhost:27017/" + dbName + "?authSource=admin"

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(dbURI).SetConnectTimeout(10 * time.Second)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Failed to disconnect from MongoDB:", err)
		}
	}()

	// Ping the database
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	log.Println("Successfully connected to MongoDB!")

	userHandler := handler.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	err = app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}
