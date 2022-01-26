package repositories

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/AndreyArthur/oganessone/src/infrastructure/models"
)

type UsersRepositoryPostgres struct {
	db *sql.DB
}

func (usersRepository *UsersRepositoryPostgres) Create(
	username string, email string, password string,
) (*entities.UserEntity, *shared.Error) {
	uuid, err := helpers.NewUuid()
	if err != nil {
		return nil, err
	}
	user, err := entities.NewUserEntity(
		uuid.Generate(),
		username,
		email,
		password,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (usersRepository *UsersRepositoryPostgres) FindByUsername(
	username string, caseSensitive bool,
) (*entities.UserEntity, *shared.Error) {
	generateQuery := func(caseSensitive bool) string {
		var field string
		if caseSensitive {
			field = "username"
		} else {
			field = "LOWER(username)"
		}
		query := strings.Join([]string{
			`SELECT 
				id, username, email, password, created_at, updated_at
			FROM
				users
			WHERE `,
			field,
			" = $1",
		}, "")
		return query
	}
	query := generateQuery(caseSensitive)
	stmt, goerr := usersRepository.db.Prepare(query)
	if goerr != nil {
		log.Println(goerr)
		return nil, exceptions.NewInternalServerError()
	}
	var queryUsername string
	if !caseSensitive {
		queryUsername = strings.ToLower(username)
	} else {
		queryUsername = username
	}
	userModel, err := models.NewUserModel()
	if err != nil {
		return nil, err
	}
	rows := stmt.QueryRow(queryUsername)
	user := userModel.Scan(rows)
	return user, nil
}

func (usersRepository *UsersRepositoryPostgres) FindByEmail(email string) (*entities.UserEntity, *shared.Error) {
	stmt, goerr := usersRepository.db.Prepare(`
		SELECT 
			id, username, email, password, created_at, updated_at
		FROM
			users
		WHERE 
			email = $1
	`)
	if goerr != nil {
		log.Println(goerr)
		return nil, exceptions.NewInternalServerError()
	}
	userModel, err := models.NewUserModel()
	if err != nil {
		return nil, err
	}
	rows := stmt.QueryRow(email)
	user := userModel.Scan(rows)
	return user, nil
}

func NewUsersRepositoryPostgres(db *sql.DB) (*UsersRepositoryPostgres, error) {
	return &UsersRepositoryPostgres{
		db: db,
	}, nil
}
