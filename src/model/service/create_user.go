package service

import (
	"fmt"

	"github.com/theus-ortiz/crud-go/src/config/logger"
	"github.com/theus-ortiz/crud-go/src/config/rest_err"
	"github.com/theus-ortiz/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init CreateUser model", zap.String("journey", "creteUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
