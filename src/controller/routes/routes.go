package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Moreno1337/api-ims/src/controller/clients"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getClientsNaturalPerson", clients.GetClientsNaturalPerson)
	r.GET("/getClientsLegalEntity", clients.GetClientsLegalEntity)
	r.POST("/registerClient", clients.RegisterClient)
	r.PUT("/updateClient/:clientId", clients.UpdateClient)
	r.DELETE("/deleteClient/:clientId", clients.DeleteClient)
}
