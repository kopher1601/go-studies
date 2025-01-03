package repositories

import (
	"errors"
	"go-gin/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
}

type memoryItemRepository struct {
	items []models.Item
}

func NewMemoryItemRepository(items []models.Item) ItemRepository {
	return &memoryItemRepository{items: items}
}

func (i *memoryItemRepository) FindAll() (*[]models.Item, error) {
	return &i.items, nil
}

func (i *memoryItemRepository) FindById(itemId uint) (*models.Item, error) {
	for _, item := range i.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("Item not found")
}

func (i *memoryItemRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(i.items) + 1)
	i.items = append(i.items, newItem)
	return &newItem, nil
}

func (i *memoryItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	for idx, item := range i.items {
		if item.ID == updateItem.ID {
			i.items[idx] = updateItem
			return &i.items[idx], nil
		}
	}
	return nil, errors.New("Unexpected error")
}

func (i *memoryItemRepository) Delete(itemId uint) error {
	for idx, item := range i.items {
		if item.ID == itemId {
			i.items = append(i.items[:idx], i.items[idx+1:]...)
			return nil
		}
	}
	return errors.New("Item not found")
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db: db}
}

func (i *itemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := i.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

func (i *itemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := i.db.First(&item, itemId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (i *itemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := i.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

func (i *itemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := i.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}

func (i *itemRepository) Delete(itemId uint) error {
	item, err := i.FindById(itemId)
	if err != nil {
		return err
	}

	// 論理削除がデフォルト
	result := i.db.Delete(item, itemId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
