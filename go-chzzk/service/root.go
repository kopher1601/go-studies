package service

import (
	"go-chzzk/repository"
	"go-chzzk/types/schema"
	"log"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) InsertChatting(user, message, roomName string) error {
	if err := s.repository.InsertChatting(user, message, roomName); err != nil {
		log.Println("Failed to chat", err)
		return err
	}
	return nil
}

func (s *Service) EnterRoom(roomName string) ([]*schema.Chat, error) {
	if res, err := s.repository.GetChatList(roomName); err != nil {
		log.Println("Failed to get all chat list", "err", err.Error())
		return nil, err
	} else {
		return res, nil
	}
}

func (s Service) RoomList() ([]*schema.Room, error) {
	if res, err := s.repository.RoomList(); err != nil {
		log.Println("Failed to get all room list", "err", err.Error())
		return nil, err
	} else {
		return res, nil
	}
}

func (s Service) MakeRoom(name string) error {
	if err := s.repository.MakeRoom(name); err != nil {
		log.Println("Failed to make a room list", "err", err.Error())
		return err
	} else {
		return nil
	}
}

func (s Service) Room(name string) (*schema.Room, error) {
	if res, err := s.repository.Room(name); err != nil {
		log.Println("Failed to get a room", "err", err.Error())
		return nil, err
	} else {
		return res, nil
	}
}
