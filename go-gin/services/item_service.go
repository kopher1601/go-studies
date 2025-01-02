package services

import (
	"go-gin/models"
	"go-gin/repositories"
)

type ItemService interface {
	FindAll() (*[]models.Item, error)
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
