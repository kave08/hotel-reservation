package main

import (
	"flag"

	"github.com/gofiber/fiber/v3"
	"github.com/kave08/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiv1.Get("/user", api.HandleGetUser)

	err := app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}

func handleFoo(c fiber.Ctx) error {

	return c.JSON(map[string]string{
		"mesg": "this is work",
	})
}
