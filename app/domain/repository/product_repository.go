package repository

import "gokomodo-assesment/app/domain/entity"

type ProductRepository interface {
	GetListProductBySellerID(sellerID string) ([]*entity.Product, error)
	GetListProduct() ([]*entity.Product, error)
	AddProduct(request *entity.Product) error
}
