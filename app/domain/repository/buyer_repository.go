package repository

import "gokomodo-assesment/app/domain/entity"

type BuyerRepository interface {
	GetBuyer(request *entity.Buyer) (*entity.Buyer, error)
}
