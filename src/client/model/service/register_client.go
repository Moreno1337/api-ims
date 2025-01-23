package service

import (
	"github.com/Moreno1337/api-ims/src/client/model"
	rest_err "github.com/Moreno1337/api-ims/src/configuration/err"
	"github.com/Moreno1337/api-ims/src/configuration/logger"
	"go.uber.org/zap"
)

func (cd *clientDomainService) RegisterClient(model.ClientDomainInterface) *rest_err.RestErr {
	journey := zap.String("journey", "RegisterClient")
	logger.Info("Init RegisterClient model", journey)

	return nil
}
