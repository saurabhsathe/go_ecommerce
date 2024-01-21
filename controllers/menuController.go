package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saurabhsathe/restaurant_ecommerce/database"
	"github.com/saurabhsathe/restaurant_ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menuCollection")

func CreateMenu(c *gin.Context) {

}
func GetAllMenus(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	result, err := menuCollection.Find(context.TODO(), bson.M{})
	defer cancel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing the menu items"})
	}
	var allMenus []bson.M
	if err = result.All(ctx, &allMenus); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, allMenus)
}
func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

func GetMenu(c *gin.Context) {

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	menuId := c.Param("menu_id")

	var menu models.Menu
	err := menuCollection.FindOne(ctx, bson.M{"Id": menuId}).Decode(&menu)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while retrieving the record requested"})
		return
	}
	c.JSON(http.StatusOK, menu)

}
func UpdateMenu(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var menu models.Menu

	if err := c.BindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menuId := c.Param("menu_id")
	filter := bson.M{"manuu_id": menuId}

	var updateObj primitive.D

	if menu.Start_Date != nil && menu.End_Date != nil {
		if !inTimeSpan(*menu.Start_Date, *menu.End_Date, time.Now()) {
			msg := "kindly retype the time"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			defer cancel()
			return
		}

		updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}
		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Category})
		}

		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := menuCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := "Menu update failed"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
