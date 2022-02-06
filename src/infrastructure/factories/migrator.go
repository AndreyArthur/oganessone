package factories

import (
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/database"
	"github.com/AndreyArthur/oganessone/src/infrastructure/helpers"
)

func MakeMigrator(environment string) (*database.Migrator, *shared.Error) {
	env, err := helpers.NewEnv()
	if err != nil {
		return nil, err
	}
	err = env.Load(environment)
	if err != nil {
		return nil, err
	}
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}
	sql, err := db.Connect()
	if err != nil {
		return nil, err
	}
	migrator, err := database.NewMigrator(sql)
	if err != nil {
		return nil, err
	}
	return migrator, nil
}
