package test_entities

import (
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

type AccountEntityTest struct{}

func (*AccountEntityTest) setup() *entities.AccountEntity {
	account, _ := entities.NewAccountEntity(&dtos.AccountDTO{
		Id:        "cc58997a-2403-af1e-7836-f0b338edcd60",
		Username:  "username",
		Email:     "account@email.com",
		Password:  "$2y$10$hRAVNUr.t6UpY1J0bQKmhO5x/K9rZPOGAPdx3HICkCrOUHR/3eyxW",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return account
}

func TestAccountEntity_isIdValid(t *testing.T) {
	// arrange
	account := (&AccountEntityTest{}).setup()
	// act
	err := account.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	account.Id = "not_an_uuid"
	// act
	err = account.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountId())
}

func TestAccountEntity_isUsernameValid(t *testing.T) {
	// arrange
	account := (&AccountEntityTest{}).setup()
	// act
	err := account.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	account.Username = "with spaces"
	// act
	err = account.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountUsername())

	// arrange
	account.Username = "toooooooooooo_big" // more than 16 chars
	// act
	err = account.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountUsername())

	// arrange
	account.Username = "sml" // less than 4 chars
	// act
	err = account.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountUsername())
}

func TestAccountEntity_isEmailValid(t *testing.T) {
	// arrange
	account := (&AccountEntityTest{}).setup()
	// act
	err := account.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	account.Email = "invalid_email"
	// act
	err = account.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountEmail())
}

func TestAccountEntity_isPasswordHashValid(t *testing.T) {
	// arrange
	account := (&AccountEntityTest{}).setup()
	// act
	err := account.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	account.Password = "not_a_bcrypt_hash"
	// act
	err = account.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPasswordHash())
}

func TestAccountEntity_IsPasswordValid(t *testing.T) {
	// arrange
	account := (&AccountEntityTest{}).setup()
	password := "p4ssword"
	// act
	err := account.IsPasswordValid(password)
	// assert
	assert.Nil(t, err)

	// arrange
	account = (&AccountEntityTest{}).setup()
	password = "password"
	// act
	err = account.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())

	// arrange
	account = (&AccountEntityTest{}).setup()
	password = "12345678"
	// act
	err = account.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())

	// arrange
	account = (&AccountEntityTest{}).setup()
	password = "to0_sml" // less than 8 characters
	// act
	err = account.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())

	// arrange
	account = (&AccountEntityTest{}).setup()
	password = "toooooooooooooooooooooooooooo_big" // more than 32 characters
	// act
	err = account.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())
}
