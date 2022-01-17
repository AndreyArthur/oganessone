package exceptions

import "github.com/AndreyArthur/murao-oganessone/src/core/shared"

const validation = "validation"

func NewInvalidUserId() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidUserId",
		"Invalid user id, must be an uuid.",
	)
}

func NewInvalidUserUsername() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidUserUsername",
		"Invalid user username, must have 4-16 characters and no whitespaces.",
	)
}

func NewInvalidUserEmail() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidUserEmail",
		"Invalid user email syntax.",
	)
}

func NewInvalidUserPassword() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidUserPassword",
		"Invalid user password, must be a bcrypt hash.",
	)
}
