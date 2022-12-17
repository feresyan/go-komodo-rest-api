package entity

type Product struct {
	BaseModel
	SellerID    string  `gorm:"column:seller_id;size:191;not null"`
	ProductName string  `gorm:"column:product_name"`
	Description string  `gorm:"column:description"`
	Price       float64 `gorm:"column:price"`
	Seller      Seller  `gorm:"foreignKey:SellerID"`
}

func (Product) TableName() string {
	return "product"
}
