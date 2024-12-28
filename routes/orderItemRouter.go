package routes

import(
	"github.com/gin-gonic/gin"
	controller "golang-restaurant-management/controllers"
)


func OrderItemsRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orderItems",controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id",controller.GetorderItem())
	incomingRoutes.GET("orderItems-order/:order_id",controller.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems",controller.CreateorderItem())
	incomingRoutes.POST("/orderItems/:orderItem_id",controller.UpdateorderItem())
}
