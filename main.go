package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":2000", "The listen address of the API server")
	flag.Parse()

	app := fiber.New()

	apiV1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)
	apiV1.Get("/user", handleUser)

	err := app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Working just fine"})
}

func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "James Foo"})
}
