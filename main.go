package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/officer47p/hotel-reservation/api"
	"github.com/officer47p/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbUrl = "mongodb://root:root@localhost:27017/"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The listen address")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbUrl))
	if err != nil {
		log.Fatal(err)
	}

	userStore := db.NewMongoUserStore(client)
	// handlers initialization
	userHandler := api.NewUserHandler(userStore)

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Get("/users/:id", userHandler.HandleGetUser)
	apiv1.Post("/users", userHandler.HandlePostUser)
	apiv1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiv1.Put("/users/:id", userHandler.HandlePutUser)
	err = app.Listen(*listenAddr)
	if err != nil {
		panic(err)
	}
}
