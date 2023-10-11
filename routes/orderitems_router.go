package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/controllers"
)

func Orderitems_routes(req *gin.Engine) {
	req.GET("/orderitems", controllers.Getsingle_orderitems())
	req.GET("/orderitems/:orderitems_id", controllers.Getsingle_orderitems())
	req.POST("/orderitems/create", controllers.Createorderitems())
	req.PATCH("/orderitems/:orderitems_id", controllers.Updateorderitems())
}
