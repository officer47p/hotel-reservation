package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 8
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (p *CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(p.FirstName) < minFirstNameLen {
		errors["firstName"] = fmt.Sprintf("first name must be at least %d", minFirstNameLen)
	}
	if len(p.LastName) < minLastNameLen {
		errors["lastName"] = fmt.Sprintf("last name must be at least %d", minLastNameLen)
	}
	if len(p.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password must be at least %d", minPasswordLen)
	}
	if !isEmailValid(p.Email) {
		errors["email"] = fmt.Sprintf("email is not valid")
	}

	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encPass, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		Password:  string(encPass),
	}, nil
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"firstName"`
	LastName  string             `bson:"lastName" json:"lastName"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"-"`
}
