package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	USER     = "SQLSERVER_USER"
	PASSWORD = "SQLSERVER_PASSWORD"
	SERVER   = "SQLSERVER_SERVER"
	DATABASE = "SQLSERVER_DATABASE"
)

type databaseParams struct {
	user     string
	password string
	server   string
	database string
}

func NewSqlServerConnection() (*sql.DB, error) {
	log.Printf("Initializing connection to SQL Server database")

	dbParams := getDatabaseConnectionParams()

	connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", dbParams.user, dbParams.password, dbParams.server, dbParams.database)
	log.Println("Connection string: " + connString)

	var err error
	var db *sql.DB

	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to SQL Server successfully!")
	return db, nil
}

func getDatabaseConnectionParams() databaseParams {
	return databaseParams{
		user:     os.Getenv(USER),
		password: os.Getenv(PASSWORD),
		server:   os.Getenv(SERVER),
		database: os.Getenv(DATABASE),
	}
}
