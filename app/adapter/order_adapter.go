package adapter

import (
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/domain/repository"
	"gorm.io/gorm"
)

type OrderAdapter struct {
	db *gorm.DB
}

func NewOrderAdapter(d *gorm.DB) repository.OrderRepository {
	return &OrderAdapter{
		db: d,
	}
}

func (a *OrderAdapter) GetOrderBySellerID(sellerID string) ([]*entity.Order, error) {
	var listOrder []*entity.Order

	err := a.db.Debug().Preload("Seller").Preload("Buyer").Preload("Product").Where("seller_id = ?", sellerID).Find(&listOrder).Error
	if err != nil {
		return nil, err
	}

	return listOrder, nil
}

func (a *OrderAdapter) AcceptOrder(request *entity.Order) error {

	err := a.db.Debug().Model(&entity.Order{}).Where(&entity.Order{
		BaseModel: entity.BaseModel{ID: request.ID},
		SellerID:  request.SellerID,
	}).Update("status", "Accept").Error
	if err != nil {
		return err
	}

	return nil
}

func (a *OrderAdapter) AddOrder(request *entity.Order) error {
	err := a.db.Debug().Create(&request).Error
	if err != nil {
		return err
	}

	return nil
}

func (a *OrderAdapter) GetOrderByBuyerID(buyerID string) ([]*entity.Order, error) {
	var listOrder []*entity.Order

	err := a.db.Debug().Preload("Seller").Preload("Buyer").Preload("Product").Where("buyer_id = ?", buyerID).Find(&listOrder).Error
	if err != nil {
		return nil, err
	}

	return listOrder, nil
}
