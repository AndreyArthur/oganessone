package definitions

import (
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateUserDTO struct {
	Username string
	Email    string
	Password string
}

type CreateUserResult = entities.UserEntity

type CreateUser interface {
	Execute(data *CreateUserDTO) (*CreateUserResult, *shared.Error)
}
