package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Getmenus() gin.HandlerFunc {
	return func(req *gin.Context) {

	}
}

func Getsingle_menu() gin.HandlerFunc {
	return func(req *gin.Context) {
		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)
		menu_id := req.Param("menus_id")

		var menu models.Menu

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": menu_id}).Decode(&menu)
		defer cancle()

		if err != nil {
			req.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		req.JSON(http.StatusOK, menu)
	}
}

func Createmenu() gin.HandlerFunc {
	return func(req *gin.Context) {
		var menu models.Menu

		var ctx, cancle = context.WithTimeout(context.Background(), 100*time.Second)

		if err := req.BindJSON(&menu); err != nil {
			req.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		validationErr := validate.Struct(menu)

		if validationErr != nil {
			req.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, insertErr := menuCollection.InsertOne(ctx, menu)

		if insertErr != nil {
			//msg := fmt.Sprintf("Menu item was not created")

			req.JSON(http.StatusInternalServerError, gin.H{"error": insertErr.Error()})

			return
		}
		defer cancle()
		req.JSON(http.StatusOK, result)
	}
}

func inTimeSpan(start *time.Time, end *time.Time, check time.Timer) bool {
	return start.After(time.Now()) && end.After(*start)
}

func Updatemenu() gin.HandlerFunc {
	return func(req *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		if err := req.BindJSON(&menu); err != nil {
			req.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menu_id := req.Param("menus_id")

		filter := bson.M{"menus_id": menu_id}

		var updateObj primitive.D

		if menu.Start_Date != nil && menu.End_Date != nil {
			if !inTimeSpan(menu.Start_Date, menu.End_Date, time.Now()) {
				msg := "kindly retype the time"
				req.JSON(http.StatusInternalServerError, msg)
				defer cancel()
				return
			}
			updateObj = append(updateObj, bson.E{"start_date", menu.Start_Date})
			updateObj = append(updateObj, bson.E{"end_date", menu.End_Date})

			if menu.Name != "" {
				updateObj = append(updateObj, bson.E{"name", menu.Name})
			}
			if menu.Categories != "" {
				updateObj = append(updateObj, bson.E{"categorites", menu.Categories})
			}

			menu.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

			updateObj = append(updateObj, bson.E{"update_at", menu.Update_at})

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
				msg := "menu update failed"
				req.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			}
			defer cancel()
			req.JSON(http.StatusOK, result)
		}
	}
}
