package controller

import "github.com/Moreno1337/api-ims/internal/clients/domain/service"

type ClientController interface {
}

type clientController struct {
	clientService service.ClientService
}

func NewClientController() (ClientController, error) {
	service, err := service.NewClientService()
	if err != nil {
		return nil, err
	}

	return &clientController{
		clientService: service,
	}, nil
}