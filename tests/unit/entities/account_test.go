package test_entities

import (
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

type UserEntityTest struct{}

func (*UserEntityTest) setup() *entities.AccountEntity {
	user, _ := entities.NewAccountEntity(&dtos.AccountDTO{
		Id:        "cc58997a-2403-af1e-7836-f0b338edcd60",
		Username:  "username",
		Email:     "user@email.com",
		Password:  "$2y$10$hRAVNUr.t6UpY1J0bQKmhO5x/K9rZPOGAPdx3HICkCrOUHR/3eyxW",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	return user
}

func TestUserEntity_isIdValid(t *testing.T) {
	// arrange
	user := (&UserEntityTest{}).setup()
	// act
	err := user.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Id = "not_an_uuid"
	// act
	err = user.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountId())
}

func TestUserEntity_isUsernameValid(t *testing.T) {
	// arrange
	user := (&UserEntityTest{}).setup()
	// act
	err := user.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Username = "with spaces"
	// act
	err = user.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountUsername())

	// arrange
	user.Username = "toooooooooooo_big" // more than 16 chars
	// act
	err = user.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountUsername())

	// arrange
	user.Username = "sml" // less than 4 chars
	// act
	err = user.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountUsername())
}

func TestUserEntity_isEmailValid(t *testing.T) {
	// arrange
	user := (&UserEntityTest{}).setup()
	// act
	err := user.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Email = "invalid_email"
	// act
	err = user.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountEmail())
}

func TestUserEntity_isPasswordHashValid(t *testing.T) {
	// arrange
	user := (&UserEntityTest{}).setup()
	// act
	err := user.IsValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Password = "not_a_bcrypt_hash"
	// act
	err = user.IsValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPasswordHash())
}

func TestUserEntity_IsPasswordValid(t *testing.T) {
	// arrange
	user := (&UserEntityTest{}).setup()
	password := "p4ssword"
	// act
	err := user.IsPasswordValid(password)
	// assert
	assert.Nil(t, err)

	// arrange
	user = (&UserEntityTest{}).setup()
	password = "password"
	// act
	err = user.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())

	// arrange
	user = (&UserEntityTest{}).setup()
	password = "12345678"
	// act
	err = user.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())

	// arrange
	user = (&UserEntityTest{}).setup()
	password = "to0_sml" // less than 8 characters
	// act
	err = user.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())

	// arrange
	user = (&UserEntityTest{}).setup()
	password = "toooooooooooooooooooooooooooo_big" // more than 32 characters
	// act
	err = user.IsPasswordValid(password)
	// assert
	assert.Equal(t, err, exceptions.NewInvalidAccountPassword())
}
