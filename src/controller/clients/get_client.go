package clients

import (
	"fmt"
	"log"

	"github.com/Moreno1337/api-ims/src/configuration/validation"
	"github.com/Moreno1337/api-ims/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func GetClientsNaturalPerson(c *gin.Context) {
	log.Println("GetClientsNaturalPerson controller initialized")

	var clientRequest request.ClientRequest

	if err := c.ShouldBindJSON(&clientRequest); err != nil {
		log.Printf("Error trying to marshall object, error=%s\n", err.Error())

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(clientRequest)
}

func GetClientsLegalEntity(c *gin.Context) {}
