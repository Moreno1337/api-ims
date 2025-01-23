package service

import (
	"github.com/Moreno1337/api-ims/src/client/model"
	rest_err "github.com/Moreno1337/api-ims/src/configuration/err"
)

func NewClientDomainService() ClientDomainService {
	return &clientDomainService{}
}

type clientDomainService struct {

}

type ClientDomainService interface {
	RegisterClient(model.ClientDomainInterface) *rest_err.RestErr
}
