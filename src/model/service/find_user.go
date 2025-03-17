package service

import (
	"github.com/theus-ortiz/crud-go/src/config/rest_err"
	"github.com/theus-ortiz/crud-go/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, 
	*rest_err.RestErr) {
	return nil, nil
}
