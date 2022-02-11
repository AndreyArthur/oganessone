package test_adapters

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/infrastructure/adapters"
	"github.com/AndreyArthur/oganessone/src/infrastructure/cache"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
)

type CacheAdapterTest struct{}

func (*CacheAdapterTest) setup() (*adapters.CacheAdapter, redis.Conn) {
	env, err := helpers.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = env.Load("test")
	if err != nil {
		log.Fatal(err)
	}
	redis, err := cache.NewRedis()
	if err != nil {
		log.Fatal(err)
	}
	connection, err := redis.Connect()
	if err != nil {
		log.Fatal(err)
	}
	cache, err := adapters.NewCacheAdapter(connection)
	if err != nil {
		log.Fatal(err)
	}
	return cache, connection
}

func TestCacheAdapter_Set(t *testing.T) {
	// arrange
	cache, connection := (&CacheAdapterTest{}).setup()
	defer connection.Close()
	defer connection.Do("FLUSHALL")
	// act
	err := cache.Set("foo", "bar", 0)
	result, _ := connection.Do("GET", "foo")
	value := fmt.Sprintf("%s", result)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, value, "bar")
}

func TestCacheAdapter_SetWithExpiration(t *testing.T) {
	// arrange
	cache, connection := (&CacheAdapterTest{}).setup()
	defer connection.Close()
	defer connection.Do("FLUSHALL")
	// act && assert
	err := cache.Set("foo", "bar", 1)
	result, _ := connection.Do("GET", "foo")
	value := fmt.Sprintf("%s", result)
	assert.Nil(t, err)
	assert.Equal(t, value, "bar")
	time.Sleep(time.Second * 1)
	result, _ = connection.Do("GET", "foo")
	assert.Nil(t, result)
}

func TestCacheAdapter_Get(t *testing.T) {
	// arrange
	cache, connection := (&CacheAdapterTest{}).setup()
	defer connection.Close()
	defer connection.Do("FLUSHALL")
	connection.Do("SET", "foo", "bar")
	// act
	value, err := cache.Get("foo")
	// assert
	assert.Nil(t, err)
	assert.Equal(t, value, "bar")
}
