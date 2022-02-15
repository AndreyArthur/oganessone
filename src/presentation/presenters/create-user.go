package presenters

import (
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/views"
)

type CreateAccountPresenter struct {
	createAccount definitions.CreateAccount
}

func (createAccountPresenter *CreateAccountPresenter) Handle(
	request *contracts.CreateAccountPresenterRequest,
) (*contracts.CreateAccountPresenterResponse, *shared.Error) {
	user, err := createAccountPresenter.createAccount.
		Execute(&definitions.CreateAccountDTO{
			Username: request.Body.Username,
			Email:    request.Body.Email,
			Password: request.Body.Password,
		})
	if err != nil {
		return nil, err
	}
	return &contracts.CreateAccountPresenterResponse{
		Body: &views.AccountView{
			Id:        user.Id,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func NewCreateAccountPresenter(
	createAccount definitions.CreateAccount,
) (*CreateAccountPresenter, *shared.Error) {
	return &CreateAccountPresenter{
		createAccount: createAccount,
	}, nil
}
