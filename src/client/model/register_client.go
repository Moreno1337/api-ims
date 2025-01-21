package model

import (
	rest_err "github.com/Moreno1337/api-ims/src/configuration/err"
	"github.com/Moreno1337/api-ims/src/configuration/logger"
	"go.uber.org/zap"
)

func (cd *ClientDomain) RegisterClient() *rest_err.RestErr {
	journey := zap.String("journey", "RegisterClient")
	logger.Info("Init RegisterClient model", journey)

	return nil
}