package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/controllers"
)

func noteRoutes(request *gin.Engine) {

	request.POST("/note/create", controllers.CreateNote)
	request.GET("/notes", controllers.GetAllNotes)
	request.GET("/note/shownote/:menu_id", controllers.GetNote)
	request.PATCH("/note/update/:menu_id", controllers.UpdateMenu)
}
