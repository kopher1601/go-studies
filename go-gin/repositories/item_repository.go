package repositories

import (
	"go-gin/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint, userId uint) error
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

func (i *itemRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	var item models.Item
	result := i.db.First(&item, "id = ? AND user_id = ?", itemId, userId)
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

func (i *itemRepository) Delete(itemId uint, userId uint) error {
	item, err := i.FindById(itemId, userId)
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
