package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/officer47p/hotel-reservation/db"
	"github.com/officer47p/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.userStore.DeleteUser(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(map[string]string{"deleted": id})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	var (
		id     = c.Params("id")
		params types.UpdateUserParams
	)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if errs := params.Validate(); len(errs) > 0 {
		return errors.Join(errs...)
	}

	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, params); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": id})
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if errs := params.Validate(); len(errs) > 0 {
		return errors.Join(errs...)
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		return err
	}
	if user == nil {
		return c.SendStatus(404)
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", users)
	return c.JSON(&users)
}
