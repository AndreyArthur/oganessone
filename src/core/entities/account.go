package entities

import (
	"regexp"
	"strings"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type AccountEntity struct {
	Id        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (account *AccountEntity) isIdValid() *shared.Error {
	regex := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$")
	if !regex.Match([]byte(account.Id)) {
		return exceptions.NewInvalidAccountId()
	}
	return nil
}

func (account *AccountEntity) isUsernameValid() *shared.Error {
	errorToReturn := exceptions.NewInvalidAccountUsername()
	usernameLength := len(account.Username)
	if usernameLength < 4 || usernameLength > 16 {
		return errorToReturn
	}
	regexChars := []string{"^.*\\s.*$"}
	regex := regexp.MustCompile(strings.Join(regexChars, ""))
	usernameHasWhitespaces := regex.Match([]byte(account.Username))
	if usernameHasWhitespaces {
		return errorToReturn
	}
	return nil
}

func (account *AccountEntity) isEmailValid() *shared.Error {
	regex := regexp.MustCompile("^[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~](\\.?[-!#$%&'*+/0-9=?A-Z^_a-z`{|}~])*@[a-zA-Z0-9](-*\\.?[a-zA-Z0-9])*\\.[a-zA-Z](-?[a-zA-Z0-9])+$")
	if !regex.Match([]byte(account.Email)) {
		return exceptions.NewInvalidAccountEmail()
	}
	return nil
}

func (account *AccountEntity) isPasswordHashValid() *shared.Error {
	regex := regexp.MustCompile(`^\$2[aby]?\$\d{1,2}\$[.\/A-Za-z0-9]{53}$`)
	if !regex.Match([]byte(account.Password)) {
		return exceptions.NewInvalidAccountPasswordHash()
	}
	return nil
}

func (account *AccountEntity) IsValid() *shared.Error {
	err := account.isIdValid()
	if err != nil {
		return err
	}
	err = account.isUsernameValid()
	if err != nil {
		return err
	}
	err = account.isEmailValid()
	if err != nil {
		return err
	}
	err = account.isPasswordHashValid()
	if err != nil {
		return err
	}
	return nil
}

func (account *AccountEntity) IsPasswordValid(password string) *shared.Error {
	if strings.TrimSpace(password) != password {
		return exceptions.NewInvalidAccountPassword()
	}
	if len(password) < 8 || len(password) > 32 {
		return exceptions.NewInvalidAccountPassword()
	}
	hasLetter := func(text string) bool {
		regex := regexp.MustCompile("^.*[a-zA-Z].*$")
		result := regex.Match([]byte(text))
		return result
	}
	hasDigit := func(text string) bool {
		regex := regexp.MustCompile(`^.*\d.*$`)
		result := regex.Match([]byte(text))
		return result
	}
	if !hasLetter(password) || !hasDigit(password) {
		return exceptions.NewInvalidAccountPassword()
	}
	return nil
}

func NewAccountEntity(data *dtos.AccountDTO) (*AccountEntity, *shared.Error) {
	account := &AccountEntity{
		Id:        data.Id,
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
	err := account.IsValid()
	if err != nil {
		return nil, err
	}
	return account, nil
}
