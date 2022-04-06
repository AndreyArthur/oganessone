package definitions

import (
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type RefreshSessionDTO struct {
	SessionKey string
}

type RefreshSessionResult struct {
	Account    *entities.AccountEntity
	SessionKey string
}

type RefreshSession interface {
	Execute(data *RefreshSessionDTO) (*RefreshSessionResult, *shared.Error)
}
