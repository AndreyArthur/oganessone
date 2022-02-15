package definitions

import (
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateAccountDTO struct {
	Username string
	Email    string
	Password string
}

type CreateAccountResult = entities.AccountEntity

type CreateAccount interface {
	Execute(data *CreateAccountDTO) (*CreateAccountResult, *shared.Error)
}
