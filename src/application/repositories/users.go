package repositories

import (
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type UsersRepository interface {
	FindByUsername(username string, caseSensitive bool) (*entities.UserEntity, *shared.Error)
	FindByEmail(email string) (*entities.UserEntity, *shared.Error)
	Create(data *dtos.UserDTO) (*entities.UserEntity, *shared.Error)
	Save(*entities.UserEntity) *shared.Error
}
