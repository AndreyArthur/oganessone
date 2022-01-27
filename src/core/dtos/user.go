package dtos

import "time"

type UserDTO struct {
	Id        string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
