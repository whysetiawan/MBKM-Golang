package controllers

import (
	"assignment2/httpserver/models"
	"assignment2/httpserver/params"
	"assignment2/httpserver/services"
	"assignment2/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrdersController interface {
	GetAllOrders(ctx *gin.Context)
	GetOrderById(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type ordersController struct {
	orderService services.OrdersService
}

func NewOrdersController(orderService services.OrdersService) *ordersController {
	return &ordersController{orderService}
}

func (c *ordersController) GetAllOrders(ctx *gin.Context) {
	var orders []models.Order

	orders, err := c.orderService.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewResponse(
			orders,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		orders,
		"Get Order Success",
	))
}

func (c *ordersController) GetOrderById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	order, err := c.orderService.GetOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		order,
		"Get order success",
	))
}

func (c *ordersController) CreateOrder(ctx *gin.Context) {
	var input params.OrderParams

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	order, err := c.orderService.CreateOrder(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		order,
		"Order Created",
	))
}

func (c *ordersController) UpdateOrder(ctx *gin.Context) {
	var input params.OrderParams

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	order, err := c.orderService.UpdateOrder(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		order,
		"Order Updated",
	))
}

func (c *ordersController) DeleteOrder(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.NewResponse(
			nil,
			err.Error(),
		))
		return
	}

	err = c.orderService.DeleteOrder(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.NewResponse(
			false,
			err.Error(),
		))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewResponse(
		true,
		"Order Deleted",
	))
}
