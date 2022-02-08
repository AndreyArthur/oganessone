package providers

import "github.com/AndreyArthur/oganessone/src/core/shared"

type SessionProvider interface {
	GenerateKey() (string, *shared.Error)
}
