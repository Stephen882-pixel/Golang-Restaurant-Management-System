package main

import (
	"golang-restaurant-management/database"
	"os"
	"golang-restaurant-management/routes"
	"github.com/gin-gonic/gin"
	"golang-restaurant-management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.openCollection(database.Clients, "food")*

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemsRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}