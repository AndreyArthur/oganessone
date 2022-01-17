package repositories

import "github.com/AndreyArthur/oganessone/src/core/entities"

type UsersRepository interface {
	FindByUsername(username string, caseSensitive bool) (*entities.UserEntity, error)
	FindByEmail(email string) (*entities.UserEntity, error)
	Create(username string, email string, password string) (*entities.UserEntity, error)
	Save(*entities.UserEntity) error
}
