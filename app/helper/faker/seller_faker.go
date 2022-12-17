package faker

import (
	"github.com/go-faker/faker/v4"
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/helper/util"
	"gorm.io/gorm"
	"log"
)

type SellerFaker struct {
	ID           string
	Email        string `faker:"email"`
	Name         string `faker:"first_name"`
	Password     string
	AlamatPickup string `faker:"oneof:Jakarta,Bandung,Surabaya,Tangerang,Madura,Bali,Padang"`
}

type ProductFaker struct {
	ID          string
	SellerID    string
	ProductName string  `faker:"word"`
	Description string  `faker:"sentence"`
	Price       float64 `faker:"oneof:10000.00,15000.00,12500.00,21000.00,65000.00"`
}

func GenerateSellerFaker(db *gorm.DB) {

	var entitySeller *entity.Seller
	err := db.Table("seller").First(&entitySeller).Error
	if err != nil {
		log.Println("find entity buyer failed")
	}

	if entitySeller.ID != "" {
		log.Println("skip seller faker")
		return
	}

	for i := 0; i < 10; i++ {
		seller := SellerFaker{}
		err = faker.FakeData(&seller)
		if err != nil {
			log.Println("failed to generate seller faker")
		}

		seller.Password, _ = util.HashPassword("123")

		err = db.Table("seller").Create(&seller).Error
		if err != nil {
			log.Println("failed to create seller")
		}

		var products []ProductFaker
		for j := 0; j < 5; j++ {
			product := ProductFaker{}
			err = faker.FakeData(&product)
			if err != nil {
				log.Println("failed to generate product faker")
			}
			product.SellerID = seller.ID
			products = append(products, product)

		}

		err = db.Table("product").Create(&products).Error
		if err != nil {
			log.Println("failed to create product")
		}
	}
}
