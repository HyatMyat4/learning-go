package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/controllers"
)

func Menu(req *gin.Engine) {
	req.GET("/menus", controllers.Getmenus())
	req.GET("/menus/:menus_id", controllers.Getsingle_menu())
	req.POST("/menus/create", controllers.Createmenu())
	req.PATCH("/menus/:menus_id", controllers.Updatemenu())
}
