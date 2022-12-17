package entity

type Seller struct {
	BaseModel
	Email        string `gorm:"column:email;unique;not null"`
	Name         string `gorm:"column:name;not null"`
	Password     string `gorm:"column:password;not null"`
	AlamatPickup string `gorm:"column:alamat_pickup;not null;type:text"`
}

func (Seller) TableName() string {
	return "seller"
}
