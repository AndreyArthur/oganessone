package entities

import (
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func setup() *UserEntity {
	user, _ := NewUserEntity(
		"cc58997a-2403-af1e-7836-f0b338edcd60",
		"username",
		"user@email.com",
		"$2y$10$hRAVNUr.t6UpY1J0bQKmhO5x/K9rZPOGAPdx3HICkCrOUHR/3eyxW",
		time.Now(),
		time.Now(),
	)
	return user
}

func TestUserEntity_isIdValid(t *testing.T) {
	// arrange
	user := setup()
	// act
	err := user.isIdValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Id = "not_an_uuid"
	// act
	err = user.isIdValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidUserId())
}

func TestUserEntity_isUsernameValid(t *testing.T) {
	// arrange
	user := setup()
	// act
	err := user.isUsernameValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Username = "with spaces"
	// act
	err = user.isUsernameValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidUserUsername())

	// arrange
	user.Username = "toooooooooooo_big" // more than 16 chars
	// act
	err = user.isUsernameValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidUserUsername())

	// arrange
	user.Username = "sml" // less than 4 chars
	// act
	err = user.isUsernameValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidUserUsername())
}

func TestUserEntity_isEmailValid(t *testing.T) {
	// arrange
	user := setup()
	// act
	err := user.isEmailValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Email = "invalid_email"
	// act
	err = user.isEmailValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidUserEmail())
}

func TestUserEntity_isPasswordValid(t *testing.T) {
	// arrange
	user := setup()
	// act
	err := user.isPasswordValid()
	// assert
	assert.Nil(t, err)

	// arrange
	user.Password = "not_a_bcrypt_hash"
	// act
	err = user.isPasswordValid()
	// assert
	assert.Equal(t, err, exceptions.NewInvalidUserPassword())
}
