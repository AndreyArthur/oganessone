package repositories

import (
	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type AccountsRepository interface {
	FindByUsername(username string, caseSensitive bool) (*entities.AccountEntity, *shared.Error)
	FindByEmail(email string) (*entities.AccountEntity, *shared.Error)
	Create(data *dtos.AccountDTO) (*entities.AccountEntity, *shared.Error)
	Save(*entities.AccountEntity) *shared.Error
}
