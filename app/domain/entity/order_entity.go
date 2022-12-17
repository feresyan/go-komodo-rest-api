package entity

type Order struct {
	BaseModel
	SellerID                   string  `gorm:"column:seller_id;size:191;not null"`
	BuyerID                    string  `gorm:"column:buyer_id;size:191;not null"`
	ProductID                  string  `gorm:"column:product_id;not null"`
	DeliverySourceAddress      string  `gorm:"column:delivery_source_address;not null"`
	DeliveryDestinationAddress string  `gorm:"column:delivery_destination_address;not null"`
	Quantity                   int64   `gorm:"column:quantity;not null"`
	Price                      float64 `gorm:"column:price;not null"`
	TotalPrice                 float64 `gorm:"column:total_price;not null"`
	Status                     string  `gorm:"column:status;not null"`
	Buyer                      Buyer   `gorm:"foreignKey:BuyerID"`
	Seller                     Seller  `gorm:"foreignKey:SellerID"`
	Product                    Product `gorm:"foreignKey:ProductID"`
}

func (Order) TableName() string {
	return "order"
}
