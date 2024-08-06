package order

import (
	"errors"
	"restaurant-order-system/internal/menu"
)

type Service interface {
	CreateOrder(userID uint, items []Item) (*Order, error)
	GetOrder(id uint) (*Order, error)
	UpdateOrderStatus(id uint, status Status) error
	ListOrders(userID uint) ([]Order, error)
}

type ServiceImpl struct {
	repo        Repository
	menuService menu.Service
}

func NewServiceImpl(repo Repository, menuService menu.Service) Service {
	return &ServiceImpl{repo: repo, menuService: menuService}
}

func (s *ServiceImpl) CreateOrder(userID uint, items []Item) (*Order, error) {
	var total float64
	for i, item := range items {
		menuItem, err := s.menuService.GetItem(item.MenuItemID)
		if err != nil {
			return nil, err
		}
		items[i].Price = menuItem.Price
		total += menuItem.Price * float64(item.Quantity)
	}

	order := &Order{
		UserID: userID,
		Status: OrderStatusPending,
		Total:  total,
		Items:  items,
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *ServiceImpl) GetOrder(id uint) (*Order, error) {
	order, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *ServiceImpl) UpdateOrderStatus(id uint, status Status) error {
	order, err := s.repo.GetById(id)
	if err != nil {
		return err
	}
	if status == OrderStatusCancelled && order.Status != OrderStatusPending {
		return errors.New("only pending orders can be cancelled")
	}

	order.Status = status
	return s.repo.Update(order)
}

func (s *ServiceImpl) ListOrders(userID uint) ([]Order, error) {
	return s.repo.List(userID)
}
