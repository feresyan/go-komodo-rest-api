package repository

import "gokomodo-assesment/app/domain/entity"

type SellerRepository interface {
	GetSeller(request *entity.Seller) (*entity.Seller, error)
}
