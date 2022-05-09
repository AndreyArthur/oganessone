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

type AccountsRepositoryPostgres struct {
	db *sql.DB
}

func (accountsRepository *AccountsRepositoryPostgres) Create(
	data *dtos.AccountDTO,
) (*entities.AccountEntity, *shared.Error) {
	generateNeededValues := func(
		values *dtos.AccountDTO,
	) (*dtos.AccountDTO, *shared.Error) {
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
		return &dtos.AccountDTO{
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
	account, err := entities.NewAccountEntity(dto)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (accountsRepository *AccountsRepositoryPostgres) FindByUsername(
	username string, caseSensitive bool,
) (*entities.AccountEntity, *shared.Error) {
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
				accounts
			WHERE `,
			field,
			" = $1",
		}, "")
		return query
	}
	query := generateQuery(caseSensitive)
	stmt, goerr := accountsRepository.db.Prepare(query)
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
	accountModel, err := models.NewAccountModel()
	if err != nil {
		return nil, err
	}
	rows := stmt.QueryRow(queryUsername)
	account := accountModel.Scan(rows)
	return account, nil
}

func (accountsRepository *AccountsRepositoryPostgres) FindByEmail(email string) (*entities.AccountEntity, *shared.Error) {
	stmt, goerr := accountsRepository.db.Prepare(`
		SELECT 
			id, username, email, password, created_at, updated_at
		FROM
			accounts
		WHERE 
			email = $1
	`)
	if goerr != nil {
		log.Println(goerr)
		return nil, exceptions.NewInternalServerError()
	}
	accountModel, err := models.NewAccountModel()
	if err != nil {
		return nil, err
	}
	rows := stmt.QueryRow(email)
	account := accountModel.Scan(rows)
	return account, nil
}

func (accountsRepository *AccountsRepositoryPostgres) FindById(id string) (*entities.AccountEntity, *shared.Error) {
	stmt, goerr := accountsRepository.db.Prepare(`
		SELECT 
			id, username, email, password, created_at, updated_at
		FROM
			accounts
		WHERE 
			id = $1
	`)
	if goerr != nil {
		log.Println(goerr)
		return nil, exceptions.NewInternalServerError()
	}
	accountModel, err := models.NewAccountModel()
	if err != nil {
		return nil, err
	}
	rows := stmt.QueryRow(id)
	account := accountModel.Scan(rows)
	return account, nil
}

func (accountRepository *AccountsRepositoryPostgres) Save(account *entities.AccountEntity) *shared.Error {
	stmt, goerr := accountRepository.db.Prepare(`
		INSERT INTO accounts	
			( id, username, email, password, created_at, updated_at )
		VALUES ( $1, $2, $3, $4, $5, $6 )
	`)
	if goerr != nil {
		log.Println(goerr)
		return exceptions.NewInternalServerError()
	}
	_, goerr = stmt.Exec(
		account.Id,
		account.Username,
		account.Email,
		account.Password,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if goerr != nil {
		log.Println(goerr)
		return exceptions.NewInternalServerError()
	}
	return nil
}

func NewAccountsRepositoryPostgres(db *sql.DB) (*AccountsRepositoryPostgres, *shared.Error) {
	return &AccountsRepositoryPostgres{
		db: db,
	}, nil
}
