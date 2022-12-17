package adapter

import (
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/domain/repository"
	"gorm.io/gorm"
)

type ProductAdapter struct {
	db *gorm.DB
}

func NewProductAdapter(d *gorm.DB) repository.ProductRepository {
	return &ProductAdapter{
		db: d,
	}
}

func (a *ProductAdapter) GetListProductBySellerID(sellerID string) ([]*entity.Product, error) {
	var listProduct []*entity.Product

	err := a.db.Debug().Where("seller_id = ?", sellerID).Find(&listProduct).Error
	if err != nil {
		return nil, err
	}

	return listProduct, nil
}

func (a *ProductAdapter) GetListProduct() ([]*entity.Product, error) {
	var listProduct []*entity.Product

	err := a.db.Debug().Find(&listProduct).Error
	if err != nil {
		return nil, err
	}

	return listProduct, nil
}

func (a *ProductAdapter) AddProduct(request *entity.Product) error {

	err := a.db.Debug().Create(&request).Error
	if err != nil {
		return err
	}

	return nil
}
