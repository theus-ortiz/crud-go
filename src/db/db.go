package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/theus-ortiz/crud-go/src/config/logger"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file", err)
		panic(err)
	}

	// Obtém as variáveis de ambiente
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Cria a string de conexão (DSN)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
			dbUser, 
			dbPassword, 
			dbHost, 
			dbPort, 
			dbName,
		)

	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Error("Error opening database connection", err)
		panic(err)
	}

	// Testa a conexão
	if err := db.Ping(); err != nil {
		logger.Error("Error pinging database", err)
		panic(err)
	}

	logger.Info("Connected to the database")

	return &Database{DB: db}
}