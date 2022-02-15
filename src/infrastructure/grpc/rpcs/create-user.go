package rpcs

import (
	"context"

	"github.com/AndreyArthur/oganessone/src/core/shared"
	"github.com/AndreyArthur/oganessone/src/infrastructure/grpc/protobuf"
	"github.com/AndreyArthur/oganessone/src/presentation/contracts"
)

type CreateUserRpc struct {
	createUserPresenter contracts.CreateAccountPresenterContract
}

func (createUserRpc *CreateUserRpc) Perform(
	ctx context.Context, request *protobuf.CreateUserRequest,
) (*protobuf.CreateUserResponse, error) {
	data := request.GetData()
	username, email, password :=
		data.GetUsername(),
		data.GetEmail(),
		data.GetPassword()
	response, err := createUserRpc.createUserPresenter.
		Handle(&contracts.CreateAccountPresenterRequest{
			Body: &contracts.CreateAccountPresenterRequestBody{
				Username: username,
				Email:    email,
				Password: password,
			},
		})
	if err != nil {
		return &protobuf.CreateUserResponse{
			Error: &protobuf.Error{
				Type:    err.Type,
				Name:    err.Name,
				Message: err.Message,
			},
			Data: nil,
		}, nil
	}
	return &protobuf.CreateUserResponse{
		Data: &protobuf.User{
			Id:        response.Body.Id,
			Username:  response.Body.Username,
			Email:     response.Body.Email,
			CreatedAt: response.Body.CreatedAt,
			UpdatedAt: response.Body.UpdatedAt,
		},
		Error: nil,
	}, nil
}

func NewCreateUserRpc(
	presenter contracts.CreateAccountPresenterContract,
) (*CreateUserRpc, *shared.Error) {
	return &CreateUserRpc{
		createUserPresenter: presenter,
	}, nil
}
