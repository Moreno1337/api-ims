package controller

import (
	"net/http"

	"github.com/Moreno1337/api-ims/src/client/controller/model/request"
	"github.com/Moreno1337/api-ims/src/client/model"
	"github.com/Moreno1337/api-ims/src/client/view"
	"github.com/Moreno1337/api-ims/src/configuration/logger"
	"github.com/Moreno1337/api-ims/src/configuration/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
var (
	ClientDomainInterface model.ClientDomainInterface
)

func (cc *clientControllerInterface) RegisterClient(c *gin.Context) {
	journey := zap.String("journey","RegisterClient")

	logger.Info("RegisterClient controller initialized", journey)

	var request request.ClientRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Error trying to marshall object", err, journey)

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewClientDomain(&request)
	
	if err := cc.service.RegisterClient(domain); err != nil {
		c.JSON(err.Code, err)
	}

	logger.Info("Client registered succesfully", journey)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}
