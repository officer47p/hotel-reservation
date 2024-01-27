package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/foo", handleFoo)

	err := app.Listen(":2000")
	if err != nil {
		panic(err)
	}
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "Working just fine"})
}
