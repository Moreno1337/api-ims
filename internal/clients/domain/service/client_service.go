package service

import "github.com/Moreno1337/api-ims/internal/clients/domain/repository"

type ClientService interface {
}

type clientService struct {
	clientRepository repository.ClientRepository
}

func NewClientService() (ClientService, error) {
	repo, err := repository.NewClientRepository()
	if err != nil {
		return nil, err
	}

	return &clientService{
		clientRepository: repo,
	}, nil
}
