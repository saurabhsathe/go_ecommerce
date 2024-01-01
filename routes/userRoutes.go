package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/controllers"
)

func UserRoutes(request *gin.Engine) {

	request.POST("/user/signup", controllers.Signup)
	request.GET("/users", controllers.GetAllUsers)
	request.POST("/user/login", controllers.Login)
	request.GET("/user/:id", controllers.GetUser)
	request.PATCH("/user/:id", controllers.UpdateUser)

}
