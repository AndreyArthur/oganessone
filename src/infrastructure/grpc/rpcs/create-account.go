package rpcs

import (
	"context"

	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/protobuf"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
)

type CreateAccountRpc struct {
	createAccountPresenter contracts.CreateAccountPresenterContract
}

func (createAccountRpc *CreateAccountRpc) Perform(
	ctx context.Context, request *protobuf.CreateAccountRequest,
) (*protobuf.CreateAccountResponse, error) {
	data := request.GetData()
	username, email, password :=
		data.GetUsername(),
		data.GetEmail(),
		data.GetPassword()
	response, err := createAccountRpc.createAccountPresenter.
		Handle(&contracts.CreateAccountPresenterRequest{
			Body: &contracts.CreateAccountPresenterRequestBody{
				Username: username,
				Email:    email,
				Password: password,
			},
		})
	if err != nil {
		return &protobuf.CreateAccountResponse{
			Error: &protobuf.Error{
				Type:    err.Type,
				Name:    err.Name,
				Message: err.Message,
			},
			Data: nil,
		}, nil
	}
	return &protobuf.CreateAccountResponse{
		Data: &protobuf.Account{
			Id:        response.Body.Id,
			Username:  response.Body.Username,
			Email:     response.Body.Email,
			CreatedAt: response.Body.CreatedAt,
			UpdatedAt: response.Body.UpdatedAt,
		},
		Error: nil,
	}, nil
}

func NewCreateAccountRpc(
	presenter contracts.CreateAccountPresenterContract,
) (*CreateAccountRpc, *shared.Error) {
	return &CreateAccountRpc{
		createAccountPresenter: presenter,
	}, nil
}
