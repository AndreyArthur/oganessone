package presenters

import (
	"time"

	"github.com/AndreyArthur/oganessone/src/application/definitions"
	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/views"
)

type RefreshSessionPresenter struct {
	refreshSession definitions.RefreshSession
}

func (refreshSessionPresenter *RefreshSessionPresenter) Handle(
	request *contracts.RefreshSessionPresenterRequest,
) (*contracts.RefreshSessionPresenterResponse, *shared.Error) {
	result, err := refreshSessionPresenter.refreshSession.
		Execute(&definitions.RefreshSessionDTO{
			SessionKey: request.Body.SessionKey,
		})
	if err != nil {
		return nil, err
	}
	return &contracts.RefreshSessionPresenterResponse{
		Body: &contracts.RefreshSessionPresenterResponseBody{
			Account: &views.AccountView{
				Id:        result.Account.Id,
				Username:  result.Account.Username,
				Email:     result.Account.Email,
				CreatedAt: result.Account.CreatedAt.Format(time.RFC3339),
				UpdatedAt: result.Account.UpdatedAt.Format(time.RFC3339),
			},
			SessionKey: result.SessionKey,
		},
	}, nil
}

func NewRefreshSessionPresenter(
	createSession definitions.RefreshSession,
) (*RefreshSessionPresenter, *shared.Error) {
	return &RefreshSessionPresenter{
		createSession,
	}, nil
}
