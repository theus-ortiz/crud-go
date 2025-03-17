package view

import (
	"github.com/theus-ortiz/crud-go/src/controller/model/response"
	"github.com/theus-ortiz/crud-go/src/model"
)

func ConvertUserDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
