package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/restaurant/go/server/database"
	"github.com/restaurant/go/server/functions"
	"github.com/restaurant/go/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

var validate = validator.New()

func Getfood() gin.HandlerFunc {
	return func(req *gin.Context) {

	}
}

func Getsingle_food() gin.HandlerFunc {
	return func(req *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)

		foodId := req.Param("food_id")

		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)

		defer cancle()

		if err != nil {
			req.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the food "})
		}
		req.JSON(http.StatusOK, food)
	}
}

func Createfood() gin.HandlerFunc {
	return func(req *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu

		var food models.Food

		if err := req.BindJSON(&food); err != nil {
			req.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		fmt.Println(food, " <-----")
		validationErr := validate.Struct(food)

		if validationErr != nil {
			req.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)
		defer cancle()

		if err != nil {
			msg := fmt.Sprintf("menu was not found ")
			req.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.Create_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()

		var num = functions.ToFixed(food.Price, 2)

		food.Price = &num

		result, insertErr := foodCollection.InsertOne(ctx, food)

		if insertErr != nil {
			msg := fmt.Sprintf("Food item was not create ")
			req.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancle()
		req.JSON(http.StatusOK, result)
	}
}

func Updatefood() gin.HandlerFunc {
	return func(req *gin.Context) {

	}
}
