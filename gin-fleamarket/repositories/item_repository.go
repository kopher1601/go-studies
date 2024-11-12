package repositories

import (
	"errors"
	"gin-fleamarket/models"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemID uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemID uint) error
}

func NewItemRepository(items []models.Item) ItemRepository {
	return &ItemMemoryRepository{items: items}
}

type ItemMemoryRepository struct {
	items []models.Item
}

func (i *ItemMemoryRepository) Delete(itemID uint) error {
	for idx, v := range i.items {
		if v.ID == itemID {
			i.items = append(i.items[:idx], i.items[idx+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func (i *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for idx, item := range i.items {
		if item.ID == updateItem.ID {
			i.items[idx] = updateItem
			return &i.items[idx], nil
		}
	}
	return nil, errors.New("unexpected error")
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

func (i *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(i.items) + 1)
	i.items = append(i.items, newItem)
	return &newItem, nil
}
