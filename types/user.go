package types

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost     = 12
	minNameLen     = 2
	minPasswordLen = 8
	minEmailLen    = 5
)

type UpdateUserParams struct {
	FirstName string `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName  string `bson:"lastName,omitempty" json:"lastName,omitempty"`
}

func (params UpdateUserParams) Validate() []error {
	errors := []error{}
	if params.FirstName != "" && len(params.FirstName) < minNameLen {
		errors = append(errors, fmt.Errorf("firstName length should be at least %d characters", minNameLen))
	}
	if params.LastName != "" && len(params.LastName) < minNameLen {
		errors = append(errors, fmt.Errorf("lastName length should be at least %d characters", minNameLen))
	}
	return errors
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() []error {
	errors := []error{}
	if len(params.FirstName) < minNameLen {
		errors = append(errors, fmt.Errorf("firstName length should be at least %d characters", minNameLen))
	}
	if len(params.LastName) < minNameLen {
		errors = append(errors, fmt.Errorf("lastName length should be at least %d characters", minNameLen))
	}
	if len(params.Password) < minPasswordLen {
		errors = append(errors, fmt.Errorf("password length should be at least %d characters", minPasswordLen))
	}
	if !isEmailValid(params.Email) {
		errors = append(errors, fmt.Errorf("email is not valid"))
	}

	return errors
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

type User struct {
	ID                string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string `bson:"firstName" json:"firstName"`
	LastName          string `bson:"lastName" json:"lastName"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encPassword),
	}, nil
}
