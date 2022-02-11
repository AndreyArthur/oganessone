package adapters

import (
	"fmt"

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

func (cache *CacheAdapter) Get(key string) (string, *shared.Error) {
	result, err := cache.connection.Do("GET", key)
	if err != nil {
		return "", exceptions.NewInternalServerError()
	}
	if result == nil {
		return "", nil
	}
	text := fmt.Sprintf("%s", result)
	return text, nil
}

func (cache *CacheAdapter) Delete(key string) *shared.Error {
	_, err := cache.connection.Do("DEL", key)
	if err != nil {
		return exceptions.NewInternalServerError()
	}
	return nil
}

func NewCacheAdapter(connection redis.Conn) (*CacheAdapter, *shared.Error) {
	return &CacheAdapter{
		connection: connection,
	}, nil
}
