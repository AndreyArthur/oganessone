package repositories

import (
	"database/sql"
	"log"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/AndreyArthur/oganessone/src/core/dtos"
	"github.com/AndreyArthur/oganessone/src/core/entities"
	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setup() (*UsersRepositoryPostgres, *sql.DB) {
	abs, _ := filepath.Abs("../../../.env.test")
	goerr := godotenv.Load(abs)
	if goerr != nil {
		log.Fatal(goerr)
	}
	db, _ := database.NewDatabase()
	sql, _ := db.Connect()
	repo, _ := NewUsersRepositoryPostgres(sql)
	return repo, sql
}

func TestUsersRepositoryPostgres_Create(t *testing.T) {
	// arrange
	repo, _ := setup()
	username, email, password := "username", "user@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	// act
	user, err := repo.Create(username, email, password)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user.IsValid())
}

func TestUsersRepositoryPostgres_FindByUsernameCaseSensitive(t *testing.T) {
	// arrange
	repo, sql := setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password := uuid.Generate(), "username", "user@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	stmt, goerr := sql.Prepare(`
		INSERT INTO users (
			id,
			username,
			email,
			password,
			created_at,
			updated_at	
		) VALUES ( $1, $2, $3, $4, $5, $6 );
	`)
	defer sql.Query("DELETE FROM users;")
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
	user, err := repo.FindByUsername(username, true)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user.IsValid())
}

func TestUsersRepositoryPostgres_FindByUsernameCaseInsensitive(t *testing.T) {
	// arrange
	repo, sql := setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password := uuid.Generate(), "username", "user@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	stmt, goerr := sql.Prepare(`
		INSERT INTO users (
			id,
			username,
			email,
			password,
			created_at,
			updated_at	
		) VALUES ( $1, $2, $3, $4, $5, $6 );
	`)
	defer sql.Query("DELETE FROM users;")
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
	user, err := repo.FindByUsername(strings.ToUpper(username), false)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user.IsValid())
}

func TestUsersRepositoryPostgres_FindByUsernameReturnNil(t *testing.T) {
	// arrange
	repo, _ := setup()
	username := "username"
	// act
	user, err := repo.FindByUsername(username, true)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user)
}

func TestUsersRepositoryPostgres_FindByEmail(t *testing.T) {
	// arrange
	repo, sql := setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password := uuid.Generate(), "username", "user@email.com", "$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW"
	stmt, goerr := sql.Prepare(`
		INSERT INTO users (
			id,
			username,
			email,
			password,
			created_at,
			updated_at	
		) VALUES ( $1, $2, $3, $4, $5, $6 );
	`)
	defer sql.Query("DELETE FROM users;")
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
	user, err := repo.FindByEmail(email)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user.IsValid())
}

func TestUsersRepositoryPostgres_FindByEmailReturnNil(t *testing.T) {
	// arrange
	repo, _ := setup()
	email := "user@email.com"
	// act
	user, err := repo.FindByEmail(email)
	// assert
	assert.Nil(t, err)
	assert.Nil(t, user)
}

func TestUsersRepositoryPostgres_Save(t *testing.T) {
	// arrange
	repo, sql := setup()
	uuid, _ := helpers.NewUuid()
	id, username, email, password, createdAt, updatedAt :=
		uuid.Generate(),
		"username",
		"user@email.com",
		"$2a$10$KtwHGGRiKWRDEq/g/2RAguaqIqU7iJNM11aFeqcwzDhuv9jDY35uW",
		time.Now().UTC(),
		time.Now().UTC()
	user, _ := entities.NewUserEntity(&dtos.UserDTO{
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
	err := repo.Save(user)
	stmt, goerr := sql.Prepare(`
		SELECT 
			id, username, email, password, created_at, updated_at
		FROM
			users
		WHERE
			id = $1;
	`)
	if goerr != nil {
		log.Fatal(goerr)
		return
	}
	stmt.QueryRow(user.Id).Scan(
		&data.id,
		&data.username,
		&data.email,
		&data.password,
		&data.createdAt,
		&data.updatedAt,
	)
	defer sql.Query("DELETE FROM users;")
	//log.Fatal("\n", data.createdAt, "\n", user.CreatedAt)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, user.Id, data.id)
	assert.Equal(t, user.Username, data.username)
	assert.Equal(t, user.Email, data.email)
	assert.Equal(t, user.Password, data.password)
	assert.Equal(t, user.CreatedAt.Format(time.RFC3339), data.createdAt.Format(time.RFC3339))
	assert.Equal(t, user.UpdatedAt.Format(time.RFC3339), data.updatedAt.Format(time.RFC3339))
}
