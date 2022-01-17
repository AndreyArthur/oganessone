package entities

import (
	"regexp"
	"strings"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type UserEntity struct {
	Id        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *UserEntity) isIdValid() *shared.Error {
	regex := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$")
	if !regex.Match([]byte(user.Id)) {
		return exceptions.NewInvalidUserId()
	}
	return nil
}

func (user *UserEntity) isUsernameValid() *shared.Error {
	errorToReturn := exceptions.NewInvalidUserUsername()
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

func (user *UserEntity) isEmailValid() *shared.Error {
	regex := regexp.MustCompile("^[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~](\\.?[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~])*@[a-zA-Z0-9](-*\\.?[a-zA-Z0-9])*\\.[a-zA-Z](-?[a-zA-Z0-9])+$")
	if !regex.Match([]byte(user.Email)) {
		return exceptions.NewInvalidUserEmail()
	}
	return nil
}

func (user *UserEntity) isPasswordValid() *shared.Error {
	regex := regexp.MustCompile(`^\$2[aby]?\$\d{1,2}\$[.\/A-Za-z0-9]{53}$`)
	if !regex.Match([]byte(user.Password)) {
		return exceptions.NewInvalidUserPassword()
	}
	return nil
}

func (user *UserEntity) IsValid() *shared.Error {
	err := user.isIdValid()
	if err != nil {
		return err
	}
	err = user.isUsernameValid()
	if err != nil {
		return err
	}
	err = user.isEmailValid()
	if err != nil {
		return err
	}
	err = user.isPasswordValid()
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
) (*UserEntity, *shared.Error) {
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
