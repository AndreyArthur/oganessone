package rpcs

import (
	"context"

	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/protobuf"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
	"github.com/AndreyArthur/oganessone/src/presentation/presenters"
)

type CreateSessionRpc struct {
	createSessionPresenter *presenters.CreateSessionPresenter
}

func (createSessionRpc *CreateSessionRpc) Perform(
	ctx context.Context, request *protobuf.CreateSessionRequest,
) (*protobuf.CreateSessionResponse, error) {
	data := request.GetData()
	login, password := data.GetLogin(), data.GetPassword()
	response, err := createSessionRpc.createSessionPresenter.Handle(&contracts.CreateSessionPresenterRequest{
		Body: &contracts.CreateSessionPresenterRequestBody{
			Login:    login,
			Password: password,
		},
	})
	if err != nil {
		return &protobuf.CreateSessionResponse{
			Data: nil,
			Error: &protobuf.Error{
				Type:    err.Type,
				Name:    err.Name,
				Message: err.Message,
			},
		}, nil
	}
	return &protobuf.CreateSessionResponse{
		Data: &protobuf.CreateSessionResponseData{
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

func NewCreateSessionRpc(
	presenter *presenters.CreateSessionPresenter,
) (*CreateSessionRpc, *shared.Error) {
	return &CreateSessionRpc{
		createSessionPresenter: presenter,
	}, nil
}
