package entities

import (
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

	assert.Equal(t, "invalid user id, must be an uuid", err.Error())
}
