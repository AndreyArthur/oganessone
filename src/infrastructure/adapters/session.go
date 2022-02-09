package adapters

import (
	"time"

	"github.com/AndreyArthur/oganessone/src/application/providers"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
)

type SessionAdapter struct{}

func (*SessionAdapter) Generate(
	userId string,
) (*providers.SessionData, *shared.Error) {
	ONE_DAY := time.Hour * 24
	tomorrow := time.Now().UTC().Add(ONE_DAY)
	expiresIn := tomorrow.Format(time.RFC3339)
	str, _ := helpers.NewString()
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789(){}[]/~`!@#$%^&*;:?"
	key := str.Random(chars, 32)
	return &providers.SessionData{
		UserId:         userId,
		Key:            key,
		ExpirationDate: expiresIn,
	}, nil
}

func NewSessionAdapter() (*SessionAdapter, *shared.Error) {
	return &SessionAdapter{}, nil
}
