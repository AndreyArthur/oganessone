package adapters

import (
	"log"

	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"golang.org/x/crypto/bcrypt"
)

type EncrypterAdapter struct{}

func (encrypterAdapter *EncrypterAdapter) Hash(text string) (string, *shared.Error) {
	const BCRYPT_COST = 10
	hash, goerr := bcrypt.GenerateFromPassword([]byte(text), BCRYPT_COST)
	if goerr != nil {
		log.Fatal(goerr)
		return "", exceptions.NewInternalServerError()
	}
	return string(hash), nil
}

func (encrypterAdapter *EncrypterAdapter) Compare(text string, hash string) (bool, *shared.Error) {
	goerr := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if goerr != nil {
		return false, nil
	}
	return true, nil
}

func NewEncrypterAdapter() (*EncrypterAdapter, *shared.Error) {
	return &EncrypterAdapter{}, nil
}
