package providers

import "github.com/AndreyArthur/oganessone/src/core/shared"

type EncrypterProvider interface {
	Hash(text string) (string, *shared.Error)
	Compare(text string, hash string) (bool, *shared.Error)
}
