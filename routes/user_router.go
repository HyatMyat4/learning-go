package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/controllers"
)

func User_router(req *gin.Engine) {
	req.GET("/users", controllers.Getusers())
	req.GET("/user/:user_id", controllers.Getsingle_user())
	req.POST("user/signup", controllers.Singup())
	req.POST("user/login", controllers.Login())
}
