package test_adapters

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/infrastructure/adapters"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
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
	assert.Equal(t, reflect.TypeOf(sessionData.ExpirationTimeInSeconds).String(), "int")
}
