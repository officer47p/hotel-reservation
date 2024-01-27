package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/officer47p/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":2000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	apiV1.Get("/user", api.HandleGetUsers)
	apiV1.Get("/user/:id", api.HandleGetUser)

	err := app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}
