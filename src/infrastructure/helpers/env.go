package helpers

import (
	"strings"

	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/joho/godotenv"
)

type Env struct{}

func (env *Env) Load(environment string) *shared.Error {
	path, err := NewPath()
	if err != nil {
		return err
	}
	filename, err := path.File()
	if err != nil {
		return err
	}
	dirname, err := path.Dir(filename)
	if err != nil {
		return err
	}
	if environment == "production" {
		envFile := strings.Join([]string{dirname, "/../../../.env"}, "")
		goerr := godotenv.Load(envFile)
		if goerr != nil {
			return exceptions.NewInternalServerError()
		}
	} else {
		envFile := strings.Join([]string{dirname, "/../../../.env.test"}, "")
		goerr := godotenv.Load(envFile)
		if goerr != nil {
			return exceptions.NewInternalServerError()
		}
	}
	return nil
}

func NewEnv() (*Env, *shared.Error) {
	return &Env{}, nil
}
