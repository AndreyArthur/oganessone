package contracts

import (
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/presentation/views"
)

type CreateAccountPresenterRequestBody struct {
	Username string
	Email    string
	Password string
}

type CreateAccountPresenterRequest struct {
	Body *CreateAccountPresenterRequestBody
}

type CreateAccountPresenterResponse struct {
	Body *views.AccountView
}

type CreateAccountPresenterContract interface {
	Handle(request *CreateAccountPresenterRequest) (*CreateAccountPresenterResponse, *shared.Error)
}
