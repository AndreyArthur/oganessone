package test_repositories

import (
	"database/sql"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/AndreyArthur/oganessone/src/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

type AccountsRepositoryPostgresTest struct{}

func (*AccountsRepositoryPostgresTest) setup() (*repositories.AccountsRepositoryPostgres, *sql.DB) {
	env, err := helpers.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	err = env.Load("test")
	if err != nil {
		log.Fatal(err)
	}
	db, _ := database.NewDatabase()
	sql, _ := db.Connect()
	repo, _ := repositories.NewAccountsRepositoryPostgres(sql)
	return repo, sql
}

func TestAccountsRepositoryPostgres_CreateWithNeededValues(t *testing.T) {
	// arrange
	repo, _ := (&AccountsRepositoryPostgresTest{}).setup()
	username, email, password := "username", "account@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	// act
	account, err := repo.Create(&dtos.AccountDTO{
		Username: username,
		Email:    email,
		Password: password,
	})
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account.IsValid())
}

func TestAccountsRepositoryPostgres_CreateWithoutNeededValues(t *testing.T) {
	// arrange
	repo, _ := (&AccountsRepositoryPostgresTest{}).setup()
	username, email, password := "username", "account@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	// act
	first, firstErr := repo.Create(&dtos.AccountDTO{
		Email:    email,
		Password: password,
	})
	second, secondErr := repo.Create(&dtos.AccountDTO{
		Username: username,
		Password: password,
	})
	third, thirdErr := repo.Create(&dtos.AccountDTO{
		Username: username,
		Email:    email,
	})
	// assert
	assert.Nil(t, first)
	assert.Nil(t, second)
	assert.Nil(t, third)
	assert.Equal(t, firstErr, exceptions.NewInternalServerError())
	assert.Equal(t, secondErr, exceptions.NewInternalServerError())
	assert.Equal(t, thirdErr, exceptions.NewInternalServerError())
}

func TestAccountsRepositoryPostgres_CreateWithCustomValues(t *testing.T) {
	// arrange
	repo, _ := (&AccountsRepositoryPostgresTest{}).setup()
	uuid, _ := helpers.NewUuid()
	now := time.Now().UTC()
	id, username, email, password, createdAt, updatedAt :=
		uuid.Generate(),
		"username",
		"account@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW",
		now,
		now
	// act
	account, err := repo.Create(&dtos.AccountDTO{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	// assert
	assert.Nil(t, err)
	assert.Equal(t, account.Id, id)
	assert.Equal(t, account.CreatedAt, createdAt)
	assert.Equal(t, account.UpdatedAt, updatedAt)
}

func TestAccountsRepositoryPostgres_FindByUsernameCaseSensitive(t *testing.T) {
	// arrange
	repo, sql := (&AccountsRepositoryPostgresTest{}).setup()
	defer sql.Query("DELETE FROM accounts;")
	uuid, _ := helpers.NewUuid()
	id, username, email, password := uuid.Generate(), "username", "account@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	stmt, goerr := sql.Prepare(`
		INSERT INTO accounts (
			id,
			username,
			email,
			password,
			created_at,
			updated_at
		) VALUES ( $1, $2, $3, $4, $5, $6 );
	`)
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	_, goerr = stmt.Exec(id, username, email, password, time.Now(), time.Now())
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	// act
	account, err := repo.FindByUsername(username, true)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account.IsValid())
}

func TestAccountsRepositoryPostgres_FindByUsernameCaseInsensitive(t *testing.T) {
	// arrange
	repo, sql := (&AccountsRepositoryPostgresTest{}).setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password := uuid.Generate(), "username", "account@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	stmt, goerr := sql.Prepare(`
		INSERT INTO accounts (
			id,
			username,
			email,
			password,
			created_at,
			updated_at
		) VALUES ( $1, $2, $3, $4, $5, $6 );
	`)
	defer sql.Query("DELETE FROM accounts;")
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	_, goerr = stmt.Exec(id, username, email, password, time.Now(), time.Now())
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	// act
	account, err := repo.FindByUsername(strings.ToUpper(username), false)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account.IsValid())
}

func TestAccountsRepositoryPostgres_FindByUsernameReturnNil(t *testing.T) {
	// arrange
	repo, _ := (&AccountsRepositoryPostgresTest{}).setup()
	username := "username"
	// act
	account, err := repo.FindByUsername(username, true)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account)
}

func TestAccountsRepositoryPostgres_FindByEmail(t *testing.T) {
	// arrange
	repo, sql := (&AccountsRepositoryPostgresTest{}).setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password := uuid.Generate(), "username", "account@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	stmt, goerr := sql.Prepare(`
		INSERT INTO accounts (
			id,
			username,
			email,
			password,
			created_at,
			updated_at
		) VALUES ( $1, $2, $3, $4, $5, $6 );
	`)
	defer sql.Query("DELETE FROM accounts;")
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	_, goerr = stmt.Exec(id, username, email, password, time.Now(), time.Now())
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	// act
	account, err := repo.FindByEmail(email)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account.IsValid())
}

func TestAccountsRepositoryPostgres_FindByEmailReturnNil(t *testing.T) {
	// arrange
	repo, _ := (&AccountsRepositoryPostgresTest{}).setup()
	email := "account@email.com"
	// act
	account, err := repo.FindByEmail(email)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account)
}

func TestAccountsRepositoryPostgres_FindById(t *testing.T) {
	// arrange
	repo, sql := (&AccountsRepositoryPostgresTest{}).setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password := uuid.Generate(), "username", "account@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	stmt, goerr := sql.Prepare(`
		INSERT INTO accounts (
			id,
			username,
			email,
			password,
			created_at,
			updated_at
		) VALUES ( $1, $2, $3, $4, $5, $6 );
	`)
	defer sql.Query("DELETE FROM accounts;")
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	_, goerr = stmt.Exec(id, username, email, password, time.Now(), time.Now())
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	// act
	account, err := repo.FindById(id)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account.IsValid())
}

func TestAccountsRepositoryPostgres_FindByIdReturnNil(t *testing.T) {
	// arrange
	repo, _ := (&AccountsRepositoryPostgresTest{}).setup()
	uuid, _ := helpers.NewUuid()
	id := uuid.Generate()
	// act
	account, err := repo.FindById(id)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, account)
}

func TestAccountsRepositoryPostgres_Save(t *testing.T) {
	// arrange
	repo, sql := (&AccountsRepositoryPostgresTest{}).setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password, createdAt, updatedAt :=
		uuid.Generate(),
		"username",
		"account@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW",
		time.Now().UTC(),
		time.Now().UTC()
	account, _ := entities.NewAccountEntity(&dtos.AccountDTO{
		Id:        id,
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	})
	data := struct {
		id        string
		username  string
		email     string
		password  string
		createdAt time.Time
		updatedAt time.Time
	}{}
	// act
	err := repo.Save(account)
	stmt, goerr := sql.Prepare(`
		SELECT
			id, username, email, password, created_at, updated_at
		FROM
			accounts
		WHERE
			id = $1;
	`)
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	stmt.QueryRow(account.Id).Scan(
		&data.id,
		&data.username,
		&data.email,
		&data.password,
		&data.createdAt,
		&data.updatedAt,
	)
	defer sql.Query("DELETE FROM accounts;")
	// assert
	assert.Nil(t, err)
	assert.Equal(t, account.Id, data.id)
	assert.Equal(t, account.Username, data.username)
	assert.Equal(t, account.Email, data.email)
	assert.Equal(t, account.Password, data.password)
	assert.Equal(t, account.CreatedAt.Format(time.RFC3339), data.createdAt.Format(time.RFC3339))
	assert.Equal(t, account.UpdatedAt.Format(time.RFC3339), data.updatedAt.Format(time.RFC3339))
}
