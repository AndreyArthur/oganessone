package models

import (
	"database/sql"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type AccountModel struct{}

func (accountModel *AccountModel) Scan(rows *sql.Row) *entities.AccountEntity {
	var id string
	var username string
	var email string
	var password string
	var createdAt time.Time
	var updatedAt time.Time
	rows.Scan(
		&id,
		&username,
		&email,
		&password,
		&createdAt,
		&updatedAt,
	)
	account, err := entities.NewAccountEntity(&dtos.AccountDTO{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	if err != nil {
		return nil
	}
	return account
}

func NewAccountModel() (*AccountModel, *shared.Error) {
	return &AccountModel{}, nil
}
