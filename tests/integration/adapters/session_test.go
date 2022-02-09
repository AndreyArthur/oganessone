package test_adapters

import (
	"math/rand"
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/infrastructure/adapters"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/AndreyArthur/oganessone/tests/helpers/verifier"
	"github.com/stretchr/testify/assert"
)

func TestSessionAdapter_Genarate(t *testing.T) {
	// arrange
	rand.Seed(time.Now().UnixNano())
	session, _ := adapters.NewSessionAdapter()
	uuid, _ := helpers.NewUuid()
	id := uuid.Generate()
	// act
	sessionData, err := session.Generate(id)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, sessionData.UserId, id)
	assert.Equal(t, len(sessionData.Key), 32)
	assert.True(t, verifier.IsISO8601(sessionData.ExpirationDate))
}
