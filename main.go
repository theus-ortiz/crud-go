package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/theus-ortiz/crud-go/src/config/logger"
	"github.com/theus-ortiz/crud-go/src/controller"
	"github.com/theus-ortiz/crud-go/src/controller/routes"
	"github.com/theus-ortiz/crud-go/src/db"
	"github.com/theus-ortiz/crud-go/src/model/repository"
	"github.com/theus-ortiz/crud-go/src/model/service"
)
func initDependencies(
	database *db.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}

func main() {
	logger.Info("Starting the application...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database := db.NewDatabase()

	//Inicia Dependencias
	userController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
