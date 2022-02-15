package contracts

import (
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/presentation/views"
)

type CreateSessionPresenterRequestBody struct {
	Login    string
	Password string
}

type CreateSessionPresenterRequest struct {
	Body *CreateSessionPresenterRequestBody
}

type CreateSessionPresenterResponseBody struct {
	Account    *views.AccountView
	SessionKey string
}

type CreateSessionPresenterResponse struct {
	Body *CreateSessionPresenterResponseBody
}

type CreateSessionPresenterContract interface {
	Handle(request *CreateSessionPresenterRequest) (*CreateSessionPresenterResponse, *shared.Error)
}
