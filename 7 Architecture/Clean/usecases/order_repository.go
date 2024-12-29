package usecases

import (
	"nhongsun/entities"
)

type OrderRepository interface {
	Save(order entities.Order) error
}
