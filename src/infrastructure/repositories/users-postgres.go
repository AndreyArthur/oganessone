package repositories

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/entities"
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
	uuid, goerr := helpers.NewUuid()
	if goerr != nil {
		log.Fatal(goerr)
		return nil, nil
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
		log.Fatal(goerr)
		return nil, nil
	}
	var queryUsername string
	if !caseSensitive {
		queryUsername = strings.ToLower(username)
	} else {
		queryUsername = username
	}
	userModel, goerr := models.NewUserModel()
	if goerr != nil {
		log.Fatal(goerr)
		return nil, nil
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
