package repositories

import "go-gin/models"

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
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
