package repository

import "github.com/Moreno1337/api-ims/internal/clients/infrastructure/persistence"

type ClientRepository interface {
}

type clientRepository struct {
	clientPersistence persistence.ClientPersistence
}

func NewClientRepository() (ClientRepository, error) {
	persistence, err := persistence.NewClientPersistence()
	if err != nil {
		return nil, err
	}

	return &clientRepository{
		clientPersistence: persistence,
	}, nil
}
