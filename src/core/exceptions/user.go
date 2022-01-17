package exceptions

import "github.com/AndreyArthur/oganessone/src/core/shared"

const validation = "validation"
const conflict = "conflict"

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

func NewUserUsernameAlreadyInUse() *shared.Error {
	return shared.NewError(
		conflict,
		"UserUsernameAlreadyInUse",
		"User username is already in use.",
	)
}

func NewUserEmailAlreadyInUse() *shared.Error {
	return shared.NewError(
		conflict,
		"UserEmailAlreadyInUse",
		"User email is already in use.",
	)
}
