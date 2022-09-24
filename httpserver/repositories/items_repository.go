package repositories

import (
	"assignment2/httpserver/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetItemByItemId(ItemId int) (models.Item, error)
	CreateItem(item models.Item) (models.Item, error)
	UpdateItem(ItemId int, item models.Item) (models.Item, error)
	DeleteItem(item models.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (repo *itemRepository) GetItemByItemId(ItemId int) (models.Item, error) {
	var item models.Item

	err := repo.db.Where("item_id = ?", ItemId).First(&item).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (repo *itemRepository) CreateItem(item models.Item) (models.Item, error) {
	err := repo.db.Create(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (repo *itemRepository) UpdateItem(itemId int, item models.Item) (models.Item, error) {
	err := repo.db.Where("item_id = ?", itemId).Updates(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (repo *itemRepository) DeleteItem(item models.Item) error {
	err := repo.db.Delete(item).Error
	if err != nil {
		return err
	}

	return err
}
