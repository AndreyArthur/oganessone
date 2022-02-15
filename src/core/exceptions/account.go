package exceptions

import "github.com/AndreyArthur/oganessone/src/core/shared"

const validation = "validation"
const conflict = "conflict"
const authentication = "authentication"

func NewInvalidAccountId() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidAccountId",
		"Invalid user id, must be an uuid.",
	)
}

func NewInvalidAccountUsername() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidAccountUsername",
		"Invalid user username, must have 4-16 characters and no whitespaces.",
	)
}

func NewInvalidAccountEmail() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidAccountEmail",
		"Invalid user email syntax.",
	)
}

func NewInvalidAccountPasswordHash() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidAccountPasswordHash",
		"Invalid user password hash, must be a bcrypt.",
	)
}

func NewInvalidAccountPassword() *shared.Error {
	return shared.NewError(
		validation,
		"InvalidAccountPassword",
		"Invalid user password, must have ascii characters, numbers and 8-32 characters.",
	)
}

func NewAccountUsernameAlreadyInUse() *shared.Error {
	return shared.NewError(
		conflict,
		"AccountUsernameAlreadyInUse",
		"Account username is already in use.",
	)
}

func NewAccountEmailAlreadyInUse() *shared.Error {
	return shared.NewError(
		conflict,
		"AccountEmailAlreadyInUse",
		"Account email is already in use.",
	)
}

func NewAccountLoginFailed() *shared.Error {
	return shared.NewError(
		authentication,
		"AccountLoginFailed",
		"Login failed, invalid login/password combination.",
	)
}
