package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/controllers"
)

func MenuRoutes(request *gin.Engine) {

	request.POST("/menu/create", controllers.CreateMenu)
	request.GET("/menus", controllers.GetAllMenus)
	request.GET("/menu/showmenu/:menu_id", controllers.GetMenu)
	request.PATCH("/food/update/:menu_id", controllers.GetMenu)
}
