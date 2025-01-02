package services

import (
	"go-gin/dto"
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemId uint) error
}

type itemService struct {
	repository repositories.ItemRepository
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &itemService{repository: repository}
}

func (i *itemService) FindAll() (*[]models.Item, error) {
	return i.repository.FindAll()
}

func (i *itemService) FindById(itemId uint) (*models.Item, error) {
	return i.repository.FindById(itemId)
}

func (i *itemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return i.repository.Create(newItem)
}

func (i *itemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := i.FindById(itemId)
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

func (i *itemService) Delete(itemId uint) error {
	return i.repository.Delete(itemId)
}
