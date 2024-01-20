package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/routes"
)

func main() {

	envport := "8000"

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + envport)

}
