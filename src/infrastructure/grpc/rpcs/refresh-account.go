package rpcs

import (
	"context"

	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/protobuf"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/presenters"
)

type RefreshSessionRpc struct {
	refreshSessionPresenter *presenters.RefreshSessionPresenter
}

func (refreshSessionRpc *RefreshSessionRpc) Perform(
	ctx context.Context, request *protobuf.RefreshSessionRequest,
) (*protobuf.RefreshSessionResponse, error) {
	data := request.GetData()
	sessionKey := data.GetSessionKey()
	response, err := refreshSessionRpc.refreshSessionPresenter.Handle(&contracts.RefreshSessionPresenterRequest{
		Body: &contracts.RefreshSessionPresenterRequestBody{
			SessionKey: sessionKey,
		},
	})
	if err != nil {
		return &protobuf.RefreshSessionResponse{
			Data: nil,
			Error: &protobuf.Error{
				Type:    err.Type,
				Name:    err.Name,
				Message: err.Message,
			},
		}, nil
	}
	return &protobuf.RefreshSessionResponse{
		Data: &protobuf.RefreshSessionResponseData{
			Account: &protobuf.Account{
				Id:        response.Body.Account.Id,
				Username:  response.Body.Account.Username,
				Email:     response.Body.Account.Email,
				CreatedAt: response.Body.Account.CreatedAt,
				UpdatedAt: response.Body.Account.UpdatedAt,
			},
		},
		Error: nil,
	}, nil
}

func NewRefreshSessionRpc(
	presenter *presenters.RefreshSessionPresenter,
) (*RefreshSessionRpc, *shared.Error) {
	return &RefreshSessionRpc{
		refreshSessionPresenter: presenter,
	}, nil
}
