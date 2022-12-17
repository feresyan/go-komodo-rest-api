package adapter

import (
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/domain/repository"
	"gorm.io/gorm"
)

type SellerAdapter struct {
	db *gorm.DB
}

func NewSellerAdapter(d *gorm.DB) repository.SellerRepository {
	return &SellerAdapter{
		db: d,
	}
}

func (a *SellerAdapter) GetSeller(request *entity.Seller) (*entity.Seller, error) {
	var seller *entity.Seller

	err := a.db.Debug().Where(&entity.Seller{Email: request.Email}).First(&seller).Error
	if err != nil {
		return nil, err
	}

	return seller, nil
}
