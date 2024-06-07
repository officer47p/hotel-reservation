package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", handleRoot)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func handleRoot(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "dnfksn"})
}
