package entities

import (
	"github.com/AndreyArthur/murao-oganessone/src/core/exceptions"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

func generateUser() *UserEntity {
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
	user := generateUser()
	err := user.isIdValid()

	assert.Nil(t, err)

	user.Id = "not_an_uuid"
	err = user.isIdValid()

	assert.Equal(t, err, exceptions.NewInvalidUserId())
}

func TestUserEntity_isUsernameValid(t *testing.T) {
	user := generateUser()
	err := user.isUsernameValid()

	assert.Nil(t, err)

	user.Username = "with spaces"
	err = user.isUsernameValid()

	assert.Equal(t, err, exceptions.NewInvalidUserUsername())

	user.Username = "toooooooooooo_big" // more than 16 chars
	err = user.isUsernameValid()

	assert.Equal(t, err, exceptions.NewInvalidUserUsername())

	user.Username = "sml" // less than 4 chars
	err = user.isUsernameValid()

	assert.Equal(t, err, exceptions.NewInvalidUserUsername())
}

func TestUserEntity_isEmailValid(t *testing.T) {
	user := generateUser()
	err := user.isEmailValid()

	assert.Nil(t, err)

	user.Email = "invalid_email"
	err = user.isEmailValid()

	assert.Equal(t, err, exceptions.NewInvalidUserEmail())
}

func TestUserEntity_isPasswordValid(t *testing.T) {
	user := generateUser()
	err := user.isPasswordValid()

	assert.Nil(t, err)

	user.Password = "not_a_bcrypt_hash"
	err = user.isPasswordValid()

	assert.Equal(t, err, exceptions.NewInvalidUserPassword())
}
