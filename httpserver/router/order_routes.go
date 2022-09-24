package router

import (
	controllers "assignment2/httpserver/contollers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(route *gin.Engine, ordersController controllers.OrdersController) {
	orderRoutes := route.Group("/order")
	{
		orderRoutes.GET("", ordersController.GetAllOrders)
		orderRoutes.POST("", ordersController.CreateOrder)
		orderRoutes.GET("/:id", ordersController.GetOrderById)
	}
}
