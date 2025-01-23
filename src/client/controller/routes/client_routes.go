package clientroutes

import (
	"github.com/gin-gonic/gin"
	"github.com/Moreno1337/api-ims/src/client/controller"
)

func InitRoutes(r *gin.RouterGroup, clientController controller.ClientControllerInterface) {
	r.GET("/getClientsNaturalPerson", clientController.GetClientsNaturalPerson)
	r.GET("/getClientsLegalEntity", clientController.GetClientsLegalEntity)
	r.POST("/registerClient", clientController.RegisterClient)
	r.PUT("/updateClient/:clientId", clientController.UpdateClient)
	r.DELETE("/deleteClient/:clientId", clientController.DeleteClient)
}
