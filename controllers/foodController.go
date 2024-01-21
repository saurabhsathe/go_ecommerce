package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/saurabhsathe/restaurant_ecommerce/database"
	"github.com/saurabhsathe/restaurant_ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "foodCollection")

func CreateFood(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var food models.Food

	err := c.BindJSON(&food)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validatioErr := validate.Struct(food)
	if validatioErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validatioErr.Error()})
		fmt.Println(validatioErr)
		return
	}

	err = menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&food)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	food.ID = primitive.NewObjectID()
	food.Food_id = food.ID.Hex()
	var num = *food.Price
	food.Price = &num

	result, insertErr := foodCollection.InsertOne(ctx, food)
	if insertErr != nil {
		msg := fmt.Sprintf("Failed to create the food item")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()
	c.JSON(http.StatusOK, result)

}
func GetAllFoods(c *gin.Context) {

}

func GetFood(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var food models.Food
	foodId := c.Param("food_id")

	err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("food not found or something went wrong")
		return
	}
	c.JSON(http.StatusOK, food)

}
func UpdateFood(c *gin.Context) {

}
