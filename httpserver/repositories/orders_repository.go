package repositories

import (
	"assignment2/httpserver/models"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	GetOrders() ([]models.Order, error)
	GetOrderByOrderId(orderId int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	UpdateOrder(orderId int, order models.Order) (models.Order, error)
	DeleteOrder(order models.Order) error
}

type ordersRepository struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *ordersRepository {
	return &ordersRepository{db}
}

func (repo *ordersRepository) GetOrders() ([]models.Order, error) {
	var orders []models.Order

	err := repo.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (repo *ordersRepository) GetOrderByOrderId(orderId int) (models.Order, error) {
	var orders models.Order

	err := repo.db.Preload("Items").Where("order_id = ?", orderId).First(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (repo *ordersRepository) CreateOrder(order models.Order) (models.Order, error) {
	err := repo.db.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (repo *ordersRepository) UpdateOrder(orderId int, order models.Order) (models.Order, error) {
	err := repo.db.Preload("Items").Where("order_id = ?", orderId).Updates(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (repo *ordersRepository) DeleteOrder(order models.Order) error {
	err := repo.db.Delete(order).Error
	if err != nil {
		return err
	}

	return err
}
