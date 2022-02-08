package definitions

import (
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateSessionDTO struct {
	Login    string
	Password string
}

type CreateSessionResult struct {
	User       *entities.UserEntity
	SessionKey string
}

type CreateSession interface {
	Execute(data *CreateSessionDTO) (*CreateSessionResult, *shared.Error)
}
