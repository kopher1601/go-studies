package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemID uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemID uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemID uint) error
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &ItemServiceImpl{repository: repository}
}

type ItemServiceImpl struct {
	repository repositories.ItemRepository
}

func (i *ItemServiceImpl) Delete(itemID uint) error {
	return i.repository.Delete(itemID)
}

func (i *ItemServiceImpl) Update(itemID uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := i.FindById(itemID)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}

	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}

	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}

	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return i.repository.Update(*targetItem)
}

func (i *ItemServiceImpl) FindAll() (*[]models.Item, error) {
	return i.repository.FindAll()
}

func (i *ItemServiceImpl) FindById(itemID uint) (*models.Item, error) {
	return i.repository.FindById(itemID)
}

func (i *ItemServiceImpl) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return i.repository.Create(newItem)
}
