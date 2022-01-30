package definitions

import (
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type CreateUser interface {
	Execute(username string, email string, password string) (*entities.UserEntity, *shared.Error)
}
