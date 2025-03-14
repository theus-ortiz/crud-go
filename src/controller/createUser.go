package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theus-ortiz/crud-go/src/config/validation"
	"github.com/theus-ortiz/crud-go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userResquest request.UserRequest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		log.Printf("Error trying to marshal object: %s", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
}
