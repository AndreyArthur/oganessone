package helpers

import (
	"github.com/AndreyArthur/oganessone/src/core/shared"
	google_uuid "github.com/google/uuid"
)

type Uuid struct{}

func (uuid *Uuid) Generate() string {
	return google_uuid.NewString()
}

func NewUuid() (*Uuid, *shared.Error) {
	return &Uuid{}, nil
}
