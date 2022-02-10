package adapters

import (
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/gomodule/redigo/redis"
)

type CacheAdapter struct {
	connection redis.Conn
}

func (cache *CacheAdapter) Set(key string, value string, expirationTimeInSeconds int) *shared.Error {
	_, err := cache.connection.Do("SET", key, value)
	if err != nil {
		return exceptions.NewInternalServerError()
	}
	if expirationTimeInSeconds > 0 {
		_, err = cache.connection.Do("EXPIRE", key, expirationTimeInSeconds)
		if err != nil {
			return exceptions.NewInternalServerError()
		}
	}
	return nil
}

func NewCacheAdapter(connection redis.Conn) (*CacheAdapter, *shared.Error) {
	return &CacheAdapter{
		connection: connection,
	}, nil
}
