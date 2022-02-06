package helpers

import (
	"path/filepath"
	"runtime"

	"github.com/AndreyArthur/oganessone/src/core/exceptions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
)

type Path struct{}

func (path *Path) File() (string, *shared.Error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", exceptions.NewInternalServerError()
	}
	return filename, nil
}

func (path *Path) Dir(filename string) (string, *shared.Error) {
	return filepath.Dir(filename), nil
}

func NewPath() (*Path, *shared.Error) {
	return &Path{}, nil
}
