package repository

import (
	"database/sql"

	"github.com/theus-ortiz/crud-go/src/db"
	"github.com/theus-ortiz/crud-go/src/config/rest_err"
	"github.com/theus-ortiz/crud-go/src/model"
)

func NewUserRepository(
	database *db.Database,
) UserRepository {
	return &userRepository{
		db: database.DB,
	}
}

type userRepository struct {
	db *sql.DB // Use *sql.DB diretamente
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
