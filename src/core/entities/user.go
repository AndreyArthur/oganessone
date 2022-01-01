package entities

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type UserEntity struct {
	Id        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *UserEntity) isIdValid() error {
	regex := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$")
	if !regex.Match([]byte(user.Id)) {
		return errors.New("invalid user id, must be an uuid")
	}

	return nil
}

func (user *UserEntity) isUsernameValid() error {
	errorToReturn := errors.New("invalid username, should have 4-16 characters and no whitespaces")
	usernameLength := len(user.Username)

	if usernameLength < 4 || usernameLength > 16 {
		return errorToReturn
	}

	regexChars := []string{"^.*\\s.*$"}
	regex := regexp.MustCompile(strings.Join(regexChars, ""))
	usernameHasWhitespaces := regex.Match([]byte(user.Username))

	if usernameHasWhitespaces {
		return errorToReturn
	}

	return nil
}

func (user *UserEntity) IsValid() error {
	err := user.isIdValid()

	if err != nil {
		return err
	}

	err = user.isUsernameValid()

	if err != nil {
		return err
	}

	return nil
}

func NewUserEntity(
	id string,
	username string,
	email string,
	password string,
	createdAt time.Time,
	updatedAt time.Time,
) (*UserEntity, error) {
	user := &UserEntity{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	err := user.IsValid()

	if err != nil {
		return nil, err
	}

	return user, nil
}
