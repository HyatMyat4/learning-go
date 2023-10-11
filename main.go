package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/restaurant/go/server/database"
	"github.com/restaurant/go/server/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	router := gin.New()

	router.Use(gin.Logger())

	routes.User_router(router)

	//router.Use(middleware.Authentication())

	routes.Food_route(router)
	routes.Order_route(router)
	routes.Orderitems_routes(router)
	routes.Menu(router)
	routes.Table_routes(router)
	routes.Invoice_router(router)

	router.Run(":" + port)
}
