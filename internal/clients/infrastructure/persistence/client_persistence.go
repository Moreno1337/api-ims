package persistence

import (
	"database/sql"

	"github.com/Moreno1337/api-ims/internal/shared/database"
)

type ClientPersistence interface {
}

type clientPersistence struct {
	db *sql.DB
}

func NewClientPersistence() (ClientPersistence, error) {
	db, err := database.NewSqlServerConnection()
	if err != nil {
		return nil, err
	}

	return &clientPersistence{
		db: db,
	}, nil
}
