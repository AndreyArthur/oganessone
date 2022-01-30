package contracts

import "github.com/AndreyArthur/oganessone/src/presentation/views"

type CreateUserPresenterRequestBody struct {
	Username string
	Email    string
	Password string
}

type CreateUserPresenterRequest struct {
	Body *CreateUserPresenterRequestBody
}

type CreateUserPresenterResponse struct {
	Body *views.UserView
}
