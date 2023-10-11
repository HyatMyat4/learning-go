package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/controllers"
)

func Order_route(req *gin.Engine) {
	req.GET("/orders", controllers.Getorders())
	req.GET("/order/:order_id", controllers.Getsingle_order())
	req.POST("/order/create", controllers.Createorder())
	req.PATCH("/order/:order_id", controllers.Updateorder())
}
