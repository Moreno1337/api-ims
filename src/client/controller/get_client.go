package controller

import (
	"net/http"

	"github.com/Moreno1337/api-ims/src/configuration/logger"
	"github.com/Moreno1337/api-ims/src/configuration/validation"
	"github.com/Moreno1337/api-ims/src/client/controller/model/request"
	"github.com/Moreno1337/api-ims/src/client/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

func GetClientsNaturalPerson(c *gin.Context) {
	journey := zapcore.Field{
		Key:    "journey",
		String: "GetClientsNaturalPerson",
	}

	logger.Info("GetClientsNaturalPerson controller initialized", journey)

	var request request.ClientRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Error("Error trying to marshall object", err, journey)

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	response := response.ClientResponse{
		Id:                request.Id,
		CompanyName:       request.CompanyName,
		FantasyName:       request.FantasyName,
		Cnpj:              request.Cnpj,
		StateRegistration: request.StateRegistration,
		Sector:            request.Sector,
		Cep:               request.Cep,
		Number:            request.Number,
		Address:           request.Address,
		City:              request.City,
		Neighborhood:      request.Neighborhood,
		Complement:        request.Complement,
		State:             request.State,
		PersonName:        request.PersonName,
		DateOfBirth:       request.DateOfBirth,
		Cpf:               request.Cpf,
		Gender:            request.Gender,
		TelephoneNumber:   request.TelephoneNumber,
		CellphoneNumber:   request.CellphoneNumber,
		Email:             request.Email,
	}

	logger.Info("Records retrieved succesfully", journey)
	c.JSON(http.StatusOK, response)
}

func GetClientsLegalEntity(c *gin.Context) {}
