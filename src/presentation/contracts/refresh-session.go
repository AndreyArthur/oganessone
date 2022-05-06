package contracts

import (
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/presentation/views"
)

type RefreshSessionPresenterRequestBody struct {
	SessionKey string
}

type RefreshSessionPresenterRequest struct {
	Body *RefreshSessionPresenterRequestBody
}

type RefreshSessionPresenterResponseBody struct {
	Account    *views.AccountView
	SessionKey string
}

type RefreshSessionPresenterResponse struct {
	Body *RefreshSessionPresenterResponseBody
}

type RefreshSessionPresenterContract interface {
	Handle(request *RefreshSessionPresenterRequest) (*RefreshSessionPresenterResponse, *shared.Error)
}
