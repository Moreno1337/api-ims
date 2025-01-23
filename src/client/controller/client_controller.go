package controller

import (
	"github.com/Moreno1337/api-ims/src/client/model/service"
	"github.com/gin-gonic/gin"
)

func NewClientControllerInterface(serviceInterface service.ClientDomainService) ClientControllerInterface {
	return &clientControllerInterface{
		service: serviceInterface,
	}
}

type ClientControllerInterface interface {
	DeleteClient(c *gin.Context)
	GetClientsNaturalPerson(c *gin.Context)
	GetClientsLegalEntity(c *gin.Context)
	RegisterClient(c *gin.Context)
	UpdateClient(c *gin.Context)
}

type clientControllerInterface struct {
	service service.ClientDomainService
}
