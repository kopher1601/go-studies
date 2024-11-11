package repositories

import "gin-fleamarket/models"

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
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
