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
	var stmt *sql.Stmt
	var goerr error
	if !caseSensitive {
		stmt, goerr = usersRepository.db.Prepare(`
			SELECT
				id, username, email, password, created_at, updated_at
			FROM 
				users
			WHERE
				LOWER(username) = $1;
		`)
	} else {
		stmt, goerr = usersRepository.db.Prepare(`
			SELECT
				id, username, email, password, created_at, updated_at
			FROM 
				users
			WHERE
				username = $1;
		`)
	}
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

func NewUsersRepositoryPostgres(db *sql.DB) (*UsersRepositoryPostgres, error) {
	return &UsersRepositoryPostgres{
		db: db,
	}, nil
}
