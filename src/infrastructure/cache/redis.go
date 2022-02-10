package cache

import (
	"os"
	"strings"

	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/gomodule/redigo/redis"
)

type Redis struct{}

func (*Redis) Connect() (redis.Conn, *shared.Error) {
	port := os.Getenv("REDIS_PORT")
	pool := &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", strings.Join([]string{":", port}, ""))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
	conn := pool.Get()
	return conn, nil
}

func NewRedis() (*Redis, *shared.Error) {
	return &Redis{}, nil
}
