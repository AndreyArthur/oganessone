package exceptions

import "github.com/AndreyArthur/oganessone/src/core/shared"

const unexpected = "unexpected"

func NewInternalServerError() *shared.Error {
	return shared.NewError(
		unexpected,
		"InternalServerError",
		"An internal server error has occured, try again later.",
	)
}
