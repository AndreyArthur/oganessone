package providers

import "github.com/AndreyArthur/oganessone/src/core/shared"

type CacheProvider interface {
	Set(key string, value string, expirationTimeInSeconds int) *shared.Error
	Get(key string) (string, *shared.Error)
	Delete(key string) *shared.Error
}
