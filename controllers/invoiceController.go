package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoiceCollection")

func CreateInvoice(c *gin.Context) {

}
func GetAllInvoices(c *gin.Context) {

}

func GetInvoice(c *gin.Context) {

}
func UpdateInvoice(c *gin.Context) {

}
