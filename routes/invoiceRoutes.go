package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/controllers"
)

func InvoiceRoutes(request *gin.Engine) {

	request.POST("/invoice/create", controllers.CreateInvoice)
	request.GET("/invoices", controllers.GetAllInvoices)
	request.GET("/invoice/showinvoice/:invoice_id", controllers.GetInvoice)
	request.PATCH("/food/update/:invoice_id", controllers.UpdateInvoice)
}
