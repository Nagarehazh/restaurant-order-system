package menu

import (
	"errors"
	"gorm.io/gorm"
)

var ErrMenuItemNotFound = errors.New("menu item not found")

type Service interface {
	CreateItem(req *Item) error
	GetItem(id uint) (*Item, error)
	UpdateItem(req *Item) error
	DeleteItem(id uint) error
	ListItems() ([]Item, error)
}

type ServiceImpl struct {
	repo Repository
}

func NewServiceImpl(repo Repository) Service {
	return &ServiceImpl{repo: repo}
}

func (s *ServiceImpl) CreateItem(item *Item) error {
	return s.repo.CreateItem(item)
}

func (s *ServiceImpl) GetItem(id uint) (*Item, error) {
	item, err := s.repo.GetItem(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMenuItemNotFound
		}
		return nil, err
	}
	return item, nil
}

func (s *ServiceImpl) UpdateItem(item *Item) error {
	return s.repo.UpdateItem(item)
}

func (s *ServiceImpl) DeleteItem(id uint) error {
	return s.repo.DeleteItem(id)
}

func (s *ServiceImpl) ListItems() ([]Item, error) {
	return s.repo.ListItems()
}
