package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/controllers"
)

func TableRoutes(request *gin.Engine) {

	request.POST("/table/create", controllers.CreateTable)
	request.GET("/tables", controllers.GetAllTables)

	request.GET("/table/:id", controllers.GetTable)
}
