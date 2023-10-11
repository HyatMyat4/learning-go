package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/controllers"
)

func Food_route(req *gin.Engine) {
	req.GET("/foods", controllers.Getfood())
	req.GET("/foods/:food_id", controllers.Getsingle_food())
	req.POST("/foods/create", controllers.Createfood())
	req.PATCH("/foods/:food_id", controllers.Updatefood())
}
