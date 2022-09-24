package services

import (
	"assignment2/httpserver/models"
	"assignment2/httpserver/params"
	"assignment2/httpserver/repositories"
)

type OrdersService interface {
	GetOrders() ([]models.Order, error)
	GetOrder(orderId int) (models.Order, error)
	CreateOrder(input params.OrderParams) (models.Order, error)
	UpdateOrder(orderId int, input params.OrderParams) (models.Order, error)
	DeleteOrder(orderId int) error
}

type orderService struct {
	orderRepository repositories.OrdersRepository
	itemsRepository repositories.ItemRepository
}

func NewOrderService(orderRepository repositories.OrdersRepository, itemsRepository repositories.ItemRepository) *orderService {
	return &orderService{orderRepository, itemsRepository}
}

func (service *orderService) GetOrders() ([]models.Order, error) {
	orders, err := service.orderRepository.GetOrders()
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (service *orderService) GetOrder(orderId int) (models.Order, error) {
	order, err := service.orderRepository.GetOrderByOrderId(orderId)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (service *orderService) CreateOrder(input params.OrderParams) (models.Order, error) {
	var order models.Order
	order.CustomerName = input.CustomerName
	order.OrderedAt = input.OrderedAt

	order, err := service.orderRepository.CreateOrder(order)
	if err != nil {
		return order, err
	}

	for _, item := range input.Items {
		newItem, err := service.itemsRepository.CreateItem(models.Item{
			ItemCode:    item.ItemCode,
			Description: item.Description,
			Quantity:    item.Quantity,
			OrderId:     int(order.OrderId),
		})
		if err != nil {
			return order, err
		}
		order.Items = append(order.Items, newItem)
	}

	return order, nil
}

func (service *orderService) UpdateOrder(orderId int, input params.OrderParams) (models.Order, error) {
	order, err := service.orderRepository.GetOrderByOrderId(orderId)
	if err != nil {
		return order, err
	}

	for _, itemInput := range input.Items {
		item, err := service.itemsRepository.GetItemByItemId(itemInput.ItemId)
		if err != nil {
			return order, err
		}

		item.ItemCode = itemInput.ItemCode
		item.Description = itemInput.Description
		item.Quantity = itemInput.Quantity
		_, err = service.itemsRepository.UpdateItem(itemInput.ItemId, item)
		if err != nil {
			return order, err
		}
	}

	order.CustomerName = input.CustomerName
	order.OrderedAt = input.OrderedAt

	order, err = service.orderRepository.UpdateOrder(orderId, order)
	if err != nil {
		return order, err
	}

	order, err = service.orderRepository.GetOrderByOrderId(orderId)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (service *orderService) DeleteOrder(orderId int) error {
	order, err := service.orderRepository.GetOrderByOrderId(orderId)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		err := service.itemsRepository.DeleteItem(item)
		if err != nil {
			return err
		}
	}

	err = service.orderRepository.DeleteOrder(order)
	if err != nil {
		return err
	}

	return nil
}
