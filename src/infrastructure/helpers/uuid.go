package helpers

import google_uuid "github.com/google/uuid"

type Uuid struct{}

func (uuid *Uuid) Generate() string {
	return google_uuid.NewString()
}

func NewUuid() (*Uuid, error) {
	return &Uuid{}, nil
}
