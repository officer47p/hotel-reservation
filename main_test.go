package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/officer47p/hotel-reservation/api"
	"github.com/officer47p/hotel-reservation/db"
	"github.com/officer47p/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testDBUri = "mongodb://root:root@localhost:27017/"
	dbname    = "hotel-reservation-test"
)

type testDB struct {
	db.UserStore
}

func (tdb *testDB) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(_ *testing.T) *testDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDBUri))
	if err != nil {
		log.Fatal(err)
	}

	userStore := db.NewMongoUserStore(client, dbname)

	return &testDB{
		UserStore: userStore,
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := api.NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		Email:     "sss@gmail.com",
		FirstName: "sss",
		LastName:  "dfdjfdsk",
		Password:  "kfhdbskhdsbfkdsb",
	}

	b, err := json.Marshal(params)
	if err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("content-type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var user types.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Error(err)
	}

	if len(user.ID) == 0 {
		t.Error("expected a user id to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Error("expected encrypted password not to be included in the json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected first name %s, but got %s\n", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected last name %s, but got %s\n", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s, but got %s\n", params.Email, user.Email)
	}
}
