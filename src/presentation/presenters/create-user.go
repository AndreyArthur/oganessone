package presenters

import (
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/views"
)

type CreateUserPresenter struct {
	createUser definitions.CreateUser
}

func (createUserPresenter *CreateUserPresenter) Handle(
	request *contracts.CreateUserPresenterRequest,
) (*contracts.CreateUserPresenterResponse, *shared.Error) {
	user, err := createUserPresenter.createUser.Execute(
		request.Body.Username,
		request.Body.Email,
		request.Body.Password,
	)
	if err != nil {
		return nil, err
	}
	return &contracts.CreateUserPresenterResponse{
		Body: &views.UserView{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func NewCreateUserPresenter(
	createUser definitions.CreateUser,
) (*CreateUserPresenter, *shared.Error) {
	return &CreateUserPresenter{
		createUser: createUser,
	}, nil
}
