package models

import (
	"database/sql"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/entities"
)

type UserModel struct{}

func (userModel *UserModel) Scan(rows *sql.Row) *entities.UserEntity {
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
	user, err := entities.NewUserEntity(
		id,
		username,
		email,
		password,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return nil
	}
	return user
}

func NewUserModel() (*UserModel, error) {
	return &UserModel{}, nil
}
