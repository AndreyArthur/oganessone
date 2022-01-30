package repositories

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/dtos"
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
	data *dtos.UserDTO,
) (*entities.UserEntity, *shared.Error) {
	generateNeededValues := func(
		values *dtos.UserDTO,
	) (*dtos.UserDTO, *shared.Error) {
		if values.Username == "" || values.Email == "" || values.Password == "" {
			log.Println(errors.New("username email and password fields are required"))
			return nil, exceptions.NewInternalServerError()
		}
		var id, username, email, password string
		var createdAt, updatedAt time.Time
		uuid, err := helpers.NewUuid()
		if err != nil {
			return nil, err
		}
		if values.Id == "" {
			id = uuid.Generate()
		} else {
			id = values.Id
		}
		username = values.Username
		email = values.Email
		password = values.Password
		now := time.Now().UTC()
		if values.CreatedAt == (time.Time{}) {
			createdAt = now
		} else {
			createdAt = values.CreatedAt
		}
		if values.UpdatedAt == (time.Time{}) {
			updatedAt = now
		} else {
			updatedAt = values.UpdatedAt
		}
		return &dtos.UserDTO{
			Id:        id,
			Username:  username,
			Email:     email,
			Password:  password,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}, nil
	}
	dto, err := generateNeededValues(data)
	if err != nil {
		return nil, err
	}
	user, err := entities.NewUserEntity(dto)
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

func (usersRepository *UsersRepositoryPostgres) Save(user *entities.UserEntity) *shared.Error {
	stmt, goerr := usersRepository.db.Prepare(`
		INSERT INTO users	
			( id, username, email, password, created_at, updated_at )
		VALUES ( $1, $2, $3, $4, $5, $6 )
	`)
	if goerr != nil {
		log.Println(goerr)
		return exceptions.NewInternalServerError()
	}
	_, goerr = stmt.Exec(
		user.Id,
		user.Username,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if goerr != nil {
		log.Println(goerr)
		return exceptions.NewInternalServerError()
	}
	return nil
}

func NewUsersRepositoryPostgres(db *sql.DB) (*UsersRepositoryPostgres, *shared.Error) {
	return &UsersRepositoryPostgres{
		db: db,
	}, nil
}
