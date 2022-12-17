package adapter

import (
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/domain/repository"
	"gorm.io/gorm"
)

type BuyerAdapter struct {
	db *gorm.DB
}

func NewBuyerAdapter(d *gorm.DB) repository.BuyerRepository {
	return &BuyerAdapter{
		db: d,
	}
}

func (a *BuyerAdapter) GetBuyer(request *entity.Buyer) (*entity.Buyer, error) {

	var buyer *entity.Buyer
	err := a.db.Debug().Where(&entity.Buyer{Email: request.Email}).First(&buyer).Error
	if err != nil {
		return nil, err
	}

	return buyer, nil
}
