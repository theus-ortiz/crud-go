package service

import (
	"github.com/theus-ortiz/crud-go/src/config/rest_err"
	"github.com/theus-ortiz/crud-go/src/model"
)

func (*userDomainService) UpdateUser(
	userid string, 
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	return nil
}
