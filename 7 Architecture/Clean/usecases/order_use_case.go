package usecases

import (
	"errors"
	"nhongsun/entities"
)

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}

// bussiness logic in use case
type OrderSevice struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderUseCase {
	return &OrderSevice{repo: repo}
}

func (s *OrderSevice) CreateOrder(order entities.Order) error {
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}

	return s.repo.Save(order)
}
