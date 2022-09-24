package main

import (
	"assignment2/config"
	controllers "assignment2/httpserver/contollers"
	"assignment2/httpserver/repositories"
	"assignment2/httpserver/router"
	httpserver "assignment2/httpserver/router"
	"assignment2/httpserver/services"
)

func main() {

	db, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}
	app := httpserver.MainRoute(db)

	orderRepository := repositories.NewOrderRepo(db)
	itemRepository := repositories.NewItemRepo(db)
	orderService := services.NewOrderService(orderRepository, itemRepository)
	orderController := controllers.NewOrdersController(orderService)

	router.OrderRoutes(app, orderController)
	// app.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, "tested")
	// })

	app.Run(":3030")
}
