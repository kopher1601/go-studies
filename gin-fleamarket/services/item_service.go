package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
}

type ItemServiceImpl struct {
	repository repositories.ItemRepository
}

func NewItemService(repository repositories.ItemRepository) ItemService {
	return &ItemServiceImpl{repository: repository}
}

func (i *ItemServiceImpl) FindAll() (*[]models.Item, error) {
	return i.repository.FindAll()
}
