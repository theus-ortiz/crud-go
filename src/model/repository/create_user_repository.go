package repository

import (
	"context"
	"fmt"

	"github.com/theus-ortiz/crud-go/src/config/logger"
	"github.com/theus-ortiz/crud-go/src/config/rest_err"
	"github.com/theus-ortiz/crud-go/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository",
		zap.String("journey", "createUser"))

	// Prepara a query SQL para inserir um novo usuário
	query := `
		INSERT INTO users (email, password, name, age)
		VALUES (?, ?, ?, ?)
	`

	// Executa a query
	result, err := ur.db.ExecContext(
		context.Background(), // Contexto para controle de timeout/cancelamento
		query,
		userDomain.GetEmail(),
		userDomain.GetPassword(),
		userDomain.GetName(),
		userDomain.GetAge(),
	)
	if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(
			fmt.Sprintf("Error trying to create user: %s", err.Error()),
		)
	}

	// Obtém o ID do usuário inserido
	userID, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error trying to get last insert ID",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(
			fmt.Sprintf("Error trying to get last insert ID: %s", err.Error()),
		)
	}

	// Define o ID no domínio do usuário
	userDomain.SetID(fmt.Sprintf("%d", userID))

	logger.Info("User created successfully",
		zap.Int64("userID", userID),
		zap.String("journey", "createUser"))

	return userDomain, nil
}
