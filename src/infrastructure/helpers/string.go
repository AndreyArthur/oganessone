package helpers

import (
	"math/rand"

	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type String struct{}

func (*String) Random(chars string, length int) string {
	letterRunes := []rune(chars)
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewString() (*String, *shared.Error) {
	return &String{}, nil
}
