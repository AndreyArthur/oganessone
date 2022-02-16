package providers

import "github.com/AndreyArthur/oganessone/src/core/shared"

type SessionData struct {
	Key                     string
	AccountId               string
	ExpirationTimeInSeconds int
}

type SessionProvider interface {
	Generate(accountId string) (*SessionData, *shared.Error)
}
