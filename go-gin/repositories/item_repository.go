package repositories

import (
	"errors"
	"go-gin/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
}

type itemRepository struct {
	items []models.Item
}

func NewItemRepository(items []models.Item) ItemRepository {
	return &itemRepository{items: items}
}

func (i *itemRepository) FindAll() (*[]models.Item, error) {
	return &i.items, nil
}

func (i *itemRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range i.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (i *itemRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(i.items) + 1)
	i.items = append(i.items, newItem)
	return &newItem, nil
}
