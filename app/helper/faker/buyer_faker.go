package faker

import (
	"github.com/go-faker/faker/v4"
	"gokomodo-assesment/app/domain/entity"
	"gokomodo-assesment/app/helper/util"
	"gorm.io/gorm"
	"log"
)

type BuyerFaker struct {
	ID               string
	Email            string `faker:"email"`
	Name             string `faker:"first_name"`
	Password         string
	AlamatPengiriman string `faker:"oneof:Jakarta,Bandung,Surabaya,Tangerang,Madura,Bali,Padang"`
}

func GenerateBuyerFaker(db *gorm.DB) {

	var entityBuyer *entity.Buyer
	err := db.Table("buyer").First(&entityBuyer).Error
	if err != nil {
		log.Println("find entity buyer failed")
	}

	if entityBuyer.ID != "" {
		log.Println("skip buyer faker")
		return
	}

	for i := 0; i < 10; i++ {
		buyer := BuyerFaker{}
		err = faker.FakeData(&buyer)

		if err != nil {
			log.Println("failed to generate buyer faker")
		}

		buyer.Password, _ = util.HashPassword("123")

		err = db.Table("buyer").Create(&buyer).Error
		if err != nil {
			log.Println("failed to create buyer")
		}
	}
}
