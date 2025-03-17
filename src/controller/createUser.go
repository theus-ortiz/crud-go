package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/crud-go/src/config/logger"
	"github.com/theus-ortiz/crud-go/src/config/validation"
	"github.com/theus-ortiz/crud-go/src/controller/model/request"
	"github.com/theus-ortiz/crud-go/src/model"
	"github.com/theus-ortiz/crud-go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		logger.Error("Error trying to validation user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userResquest.Email,
		userResquest.Password,
		userResquest.Name,
		userResquest.Age,
	)
	if err := uc.service.CreateUser(domain); err != nil {
		logger.Error(
			"Error trying to call to CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertUserDomainToResponse(domain))
}
