package repository

import "gokomodo-assesment/app/domain/entity"

type OrderRepository interface {
	GetOrderBySellerID(sellerID string) ([]*entity.Order, error)
	GetOrderByBuyerID(buyerID string) ([]*entity.Order, error)
	AcceptOrder(request *entity.Order) error
	AddOrder(request *entity.Order) error
}
