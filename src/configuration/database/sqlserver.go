package sqlserver

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/Moreno1337/api-ims/src/configuration/logger"
	"go.uber.org/zap"
)

var (
	db *sql.DB

	USER = "SQLSERVER_USER"
	PASSWORD = "SQLSERVER_PASSWORD"
	SERVER = "SQLSERVER_SERVER"
	DATABASE = "SQLSERVER_DATABASE"
)

type databaseParams struct {
	user string
	password string
	server string
	database string
}

func InitConnection() {
	journey := zap.String("journey", "sqlserver.InitConnection")
	logger.Info("Initializing connection to SQL Server database", journey)

	dbParams := getDatabaseConnectionParams()

	connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s", dbParams.user, dbParams.password, dbParams.server, dbParams.database)
	logger.Info(fmt.Sprintf("Connection string: %s", connString), journey)

	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	logger.Info("Connected to SQL Server successfully!", journey)
}

func GetDB() *sql.DB {
	return db
}

func getDatabaseConnectionParams() databaseParams {
	return databaseParams{
		user: os.Getenv(USER),
		password: os.Getenv(PASSWORD),
		server: os.Getenv(SERVER),
		database: os.Getenv(DATABASE),
	}
}
