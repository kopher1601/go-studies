package repositories

import (
	"errors"
	"gin-fleamarket/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemID uint) (*models.Item, error)
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemRepository(items []models.Item) ItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (i *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &i.items, nil
}

func (i *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range i.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}
