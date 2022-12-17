package entity

type Buyer struct {
	BaseModel
	Email            string `gorm:"column:email;unique;not null"`
	Name             string `gorm:"column:name;not null"`
	Password         string `gorm:"column:password;not null"`
	AlamatPengiriman string `gorm:"column:alamat_pengiriman;not null;type:text"`
}

func (Buyer) TableName() string {
	return "buyer"
}
