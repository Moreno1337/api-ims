package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Moreno1337/api-ims/src/controller/clients"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getClientsNaturalPerson", controller.GetClientsNaturalPerson)
	r.GET("/getClientsLegalEntity", controller.GetClientsLegalEntity)
	r.POST("/registerClient", controller.RegisterClient)
	r.PUT("/updateClient/:clientId", controller.UpdateClient)
	r.DELETE("/deleteClient/:clientId", controller.DeleteClient)
}
