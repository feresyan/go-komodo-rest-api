package model

type ProductModel struct {
	ID          string  `json:"id"`
	SellerID    string  `json:"seller_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductResponseModel struct {
	ResponseCode string `json:"response_code"`
	ResponseDesc string `json:"response_desc"`
}
