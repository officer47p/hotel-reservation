package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/officer47p/hotel-reservation/api"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address")
	flag.Parse()

	app := fiber.New()

	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", api.HandleGetUsers)
	apiv1.Get("/users/:id", api.HandleGetUser)

	err := app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}
