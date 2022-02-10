package providers

import "github.com/AndreyArthur/oganessone/src/core/shared"

type SessionData struct {
	Key                     string
	UserId                  string
	ExpirationTimeInSeconds int
}

type SessionProvider interface {
	Generate(userId string) (*SessionData, *shared.Error)
}
