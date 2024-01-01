package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/controllers"
)

func FoodRoutes(request *gin.Engine) {

	request.POST("/food/create", controllers.CreateFood)
	request.GET("/foods", controllers.GetAllFoods)

	request.PATCH("/food/:food_id", controllers.UpdateFood)
}
